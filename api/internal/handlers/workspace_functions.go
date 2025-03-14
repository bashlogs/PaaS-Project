package handlers

import (
	"errors"
	"fmt"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/tools"
	log "github.com/sirupsen/logrus"
)

// func getEmail(r *http.Request) string {
//     secretKey := []byte("khadde")
//     cookie, err := r.Cookie("authToken")
//     if err != nil {
//         if errors.Is(err, http.ErrNoCookie) {
//             log.Error("Authorization token missing", err)
//             return ""
//         }
//         log.Error("Error retrieving cookie:", err)
//         return ""
//     }

//     tokenString := cookie.Value
//     if tokenString == "" {
//         log.Error("Authorization token is empty")
//         return ""
//     }

//     // Parse and validate the token
//     token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//             return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//         }
//         return secretKey, nil
//     })

//     if err != nil || !token.Valid {
//         log.Error("Invalid token")
//         return ""
//     }

//     claims, ok := token.Claims.(jwt.MapClaims)
//     if !ok {
//         log.Error("Failed to parse token claims")
//         return ""
//     }

//     email, ok := claims["email"].(string)
//     if !ok {
//         log.Error("Email claim is not a string")
//         return ""
//     }

//     return email
// }


func getNamespaceInfo(username string) (api.Namespace_info, error) {
    var namespaceInfo api.Namespace_info

    // Connect to the database
    database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Error("Database connection error: ", err)
        return namespaceInfo, fmt.Errorf("failed to connect to database: %w", err)
    }

    query := `
        SELECT
            COALESCE(st.type_name, default_st.type_name) AS type_name,
            COALESCE(st.cpu_limit, default_st.cpu_limit) AS cpu_limit,
            COALESCE(st.memory_limit, default_st.memory_limit) AS memory_limit,
            COALESCE(st.namespace_limit, default_st.namespace_limit) AS namespace_limit,
            COUNT(n.namespace_id) AS namespace_count,
            t.transaction_date,
            t.validity
        FROM
            users u
        LEFT JOIN
            subscription s ON u.user_id = s.user_id
        LEFT JOIN
            transaction t ON s.transaction_id = t.transaction_id AND t.validity >= CURRENT_DATE
        LEFT JOIN
            subscription_types st ON s.type_id = st.type_id
        LEFT JOIN
            namespace n ON u.username = n.username AND n.active = TRUE
        JOIN
            subscription_types default_st ON default_st.type_id = 1 and u.username = $1
        GROUP BY
            u.name, st.type_name, st.cpu_limit, st.memory_limit, st.namespace_limit,
            default_st.type_name, default_st.cpu_limit, default_st.memory_limit, default_st.namespace_limit,
            t.transaction_date, t.validity
        ORDER BY
            t.transaction_date DESC NULLS LAST
        LIMIT 1;
    `

    rows, err := database.DB.Query(query, username)
    if err != nil {
        log.Error("Error querying database: ", err)
        return namespaceInfo, fmt.Errorf("database query failed: %w", err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&namespaceInfo.TypeName, &namespaceInfo.CPULimit, &namespaceInfo.MemoryLimit, &namespaceInfo.NamespaceLimit, &namespaceInfo.NamespaceCount, &namespaceInfo.TransactionDate, &namespaceInfo.Validity)
        if err != nil {
            log.Error("Error scanning row: ", err)
            return namespaceInfo, fmt.Errorf("failed to scan row: %w", err)
        }

		return namespaceInfo, nil
    }

    log.Error("No workspaces found for user: ", username)
    return namespaceInfo, errors.New("no workspaces found for the specified user")
}

func createWorkspace(param *api.CreateWorkspace) (*api.Workspace, error) {
    // Connect to the database
    database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Error("Database connection error: ", err)
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Insert the new workspace
    insertQuery := `
        INSERT INTO namespace (username, namespace, last_modified, active)
        VALUES ($1, $2, CURRENT_TIMESTAMP, true);
    `
    _, err = database.DB.Exec(insertQuery, param.Username, param.Name)
    if err != nil {
        log.Error("Error inserting workspace: ", err)
        return nil, fmt.Errorf("failed to insert workspace: %w", err)
    }

    // Retrieve the newly created workspace
    selectQuery := `
        SELECT namespace_id, namespace, active 
        FROM namespace 
        WHERE username = $1 AND namespace = $2;
    `
    var workspace api.Workspace
    err = database.DB.QueryRow(selectQuery, param.Username, param.Name).
        Scan(&workspace.Id, &workspace.Name, &workspace.IsActive)
    
    if err != nil {
        log.Error("Error querying database: ", err)
        return nil, fmt.Errorf("database query failed: %w", err)
    }

    // Generate the workspace endpoint URL
    workspace.Endpoint = fmt.Sprintf("http://localhost:8080/%s/%s", param.Username, workspace.Name)

    // Log the workspace data
    fmt.Println("Workspace data: ", workspace)

    return &workspace, nil
}
