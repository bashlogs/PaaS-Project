package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/tools"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func getEmail(r *http.Request) string {
    secretKey := []byte("khadde")
    cookie, err := r.Cookie("authToken")
    if err != nil {
        if errors.Is(err, http.ErrNoCookie) {
            log.Error("Authorization token missing", err)
            return ""
        }
        log.Error("Error retrieving cookie:", err)
        return ""
    }

    tokenString := cookie.Value
    if tokenString == "" {
        log.Error("Authorization token is empty")
        return ""
    }

    // Parse and validate the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil || !token.Valid {
        log.Error("Invalid token")
        return ""
    }
    

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        log.Error("Failed to parse token claims")
        return ""
    }

    email, ok := claims["email"].(string)
    if !ok {
        log.Error("Email claim is not a string")
        return ""
    }

    return email
}

func GetWorkspaces(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetWorkspaces")
	database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Error("Database connection error: ", err)
        api.InternalErrorHandler(w)
        return
    }

	// check for the invalid request
	// 	select namespace, active from namespace where username = 'alicej';	
	Email := getEmail(r)
    fmt.Println("Email: ", Email)
	rows, err := database.DB.Query("select n.namespace_id, n.namespace, n.active from namespace n join users u on u.username = n.username where u.email = $1", Email)

	if err != nil {
        api.RequestErrorHandler(w, errors.New("no workspaces found"))
		return
	}

    var workspaces []api.Workspace
    for rows.Next() {
        var workspace api.Workspace
        err := rows.Scan(&workspace.Id, &workspace.Name, &workspace.IsActive)
        if err != nil {
            log.Error("Error scanning row:", err)
            api.InternalErrorHandler(w)
            return
        }
        workspace.Endpoint = fmt.Sprintf("http://localhost:8080/%s", workspace.Name)
        workspaces = append(workspaces, workspace)
    }

    if err = rows.Err(); err != nil {
        log.Error("Error iterating rows:", err)
        api.InternalErrorHandler(w)
        return
    }

    fmt.Println("Workspace data: ", workspaces)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workspaces)
	if err != nil {
		log.Error("Error encoding dashboard response:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
}


// func CreateWorkspace(w http.ResponseWriter, r *http.Request) {
//     var param = api.CreateWorkspace{}
// 	err := json.NewDecoder(r.Body).Decode(&param)

// 	if err != nil{
// 		log.Error(err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}

//     database, err := tools.ConnectToDatabase()
//     if err != nil {
//         log.Error("Database connection error: ", err)
//         api.InternalErrorHandler(w)
//         return
//     }

//     // Check if the users can create namespace or not

//     // 
//     rows, err := database.DB.Query(`
//         SELECT
//             COALESCE(st.type_name, default_st.type_name) AS type_name,
//             COALESCE(st.cpu_limit, default_st.cpu_limit) AS cpu_limit,
//             COALESCE(st.memory_limit, default_st.memory_limit) AS memory_limit,
//             COALESCE(st.namespace_limit, default_st.namespace_limit) AS namespace_limit,
//             COUNT(n.namespace_id) AS namespace_count,
//             t.transaction_date,
//             t.validity
//         FROM
//             users u
//         LEFT JOIN
//             subscription s ON u.user_id = s.user_id
//         LEFT JOIN
//             transaction t ON s.transaction_id = t.transaction_id AND t.validity >= CURRENT_DATE
//         LEFT JOIN
//             subscription_types st ON s.type_id = st.type_id
//         LEFT JOIN
//             namespace n ON u.username = n.username AND n.active = TRUE -- Count only active namespaces
//         JOIN
//             subscription_types default_st ON default_st.type_id = 1 and u.username = $1
//         GROUP BY
//             u.name, st.type_name, st.cpu_limit, st.memory_limit, st.namespace_limit,
//             default_st.type_name, default_st.cpu_limit, default_st.memory_limit, default_st.namespace_limit,
//             t.transaction_date, t.validity
//         ORDER BY
//             t.transaction_date DESC NULLS LAST;
//     `, &param.Username)

//     // for rows.Next() {
//     //     var workspace api.Workspace
//     //     err := rows.Scan(&workspace.Id, &workspace.Name, &workspace.IsActive)
//     //     if err != nil {
//     //         log.Error("Error scanning row:", err)
//     //         api.InternalErrorHandler(w)
//     //         return
//     //     }
//     //     workspace.Endpoint = fmt.Sprintf("http://localhost:8080/%s", workspace.Name)
//     //     workspaces = append(workspaces, workspace)
//     // }
    
//     if err != nil {
//         api.RequestErrorHandler(w, errors.New("no workspaces found"))
//         return
//     }

//     defer rows.Close()

//     var username string
//     for rows.Next() {
//         err := rows.Scan(&username)
//         if err != nil {
//             log.Error("Error scanning row:", err)
//             api.InternalErrorHandler(w)
//             return
//         }
//     }

//     if err = rows.Err(); err != nil {
//         log.Error("Error iterating rows:", err)
//         api.InternalErrorHandler(w)
//         return
//     }

//     var workspace api.Workspace
//     err = json.NewDecoder(r.Body).Decode(&workspace)
//     if err != nil {
//         log.Error("Error decoding request body:", err)
//         api.RequestErrorHandler(w, errors.New("invalid request"))
//         return
//     }

//     if workspace.Name == "" {
//         api.RequestErrorHandler(w, errors.New("workspace name is required"))
//         return
//     }

//     _, err = database.DB.Exec("insert into namespace (namespace, username) values ($1, $2)", workspace.Name, username)
//     if err != nil {
//         log.Error("Error inserting workspace:", err)
//         api.InternalErrorHandler(w)
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
// }