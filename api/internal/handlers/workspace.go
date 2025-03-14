package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetWorkspaces(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetWorkspaces")

	email, ok := r.Context().Value("email").(string)
	if !ok || email == "" {
		http.Error(w, "Failed to retrieve user information", http.StatusUnauthorized)
		return
	}

    database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Error("Database connection error: ", err)
        api.InternalErrorHandler(w)
        return
    }

	rows, err := database.DB.Query("select n.namespace_id, n.namespace, n.active from namespace n join users u on u.username = n.username where u.email = $1", email)

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

func CreateWorkspace(w http.ResponseWriter, r *http.Request) {
    var param = api.CreateWorkspace{}
	err := json.NewDecoder(r.Body).Decode(&param)

	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

    var namespace_info api.Namespace_info
    namespace_info, err = getNamespaceInfo(param.Username)
    
    if err != nil {
        api.RequestErrorHandler(w, err)
        return
    }

    if namespace_info.NamespaceCount >= namespace_info.NamespaceLimit {
        api.RequestErrorHandler(w, errors.New("you have reached the limit of creating workspace"))
        return
    }

    var output *api.Workspace
    output, err = createWorkspace(&param)
    
    if err != nil {
        api.InternalErrorHandler(w)
        return
    }

    response := api.Workspace {
        Id : output.Id,
        Name : output.Name,
        IsActive : output.IsActive,
        Endpoint : output.Endpoint,
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Error("Error encoding dashboard response:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
}

func DeleteWorkspace(w http.ResponseWriter, r *http.Request) {
	// Get workspace ID from URL query parameters
	workspaceID := r.URL.Query().Get("id")
	if workspaceID == "" {
		http.Error(w, "Missing workspace ID", http.StatusBadRequest)
		return
	}

    database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Error("Database connection error: ", err)
        api.InternalErrorHandler(w)
        return
    }

	// Delete workspace from the database
	query := `DELETE FROM namespace WHERE namespace_id = $1`
	result, err := database.DB.Exec(query, workspaceID)
	if err != nil {
		http.Error(w, "Failed to delete workspace", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := map[string]string{}
	if rowsAffected > 0 {
		response["message"] = "Workspace deleted successfully."
	} else {
		response["message"] = "Workspace not found."
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}