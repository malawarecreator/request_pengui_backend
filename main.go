package main
import "log"
import "net/http"
import "fmt"
import "os/exec"
func main() {

	http.HandleFunc("/api", handle_exec)
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		fmt.Println(err.Error())
	}
}


func handle_exec(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	url := r.URL.Query().Get("url")

	if command == "" || url == "" {
		fmt.Fprintf(w, "Error: missing params")
		return
	}
	if command == "TRACE" {
		cmd := exec.Command("curl", "-X", "TRACE", url)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Fprintln(w, string(out))
		return
	}
	if command == "OPTIONS" {
		cmd := exec.Command("curl", "-X", "OPTIONS", url, "-i")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
			return 
		}

		fmt.Fprintln(w, string(out))
		return
	}


}