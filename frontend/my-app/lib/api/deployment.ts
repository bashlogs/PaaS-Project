
export async function deployment(id: number, frontendUrl: string, backendUrl: string) {
    const response = await fetch(`http://127.0.0.1:8000/deployment`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ id, frontendUrl, backendUrl }),
      });
    
      if (!response.ok) {
        throw new Error("Failed to get any deployment updates");
      }
}
