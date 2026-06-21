# 📚 Τι Πρέπει να Μάθουμε — ascii-art-web
> Επιγραμματική λίστα. Για κάθε θέμα υπάρχει αναλυτικό μάθημα στο `lessons.md`.

---

## 🌐 1. Βασικές Έννοιες Web (HTTP)
- [ ] Τι είναι HTTP/HTTPS και πώς λειτουργεί το request–response cycle
- [ ] Client–Server μοντέλο: browser = client, Go program = server
- [ ] HTTP Methods: **GET** (ζητώ δεδομένα) vs **POST** (στέλνω δεδομένα)
- [ ] HTTP Status Codes: 200, 400, 404, 500 και τι σημαίνει το καθένα
- [ ] Τι είναι URL, path, endpoint, query string
- [ ] Τι είναι HTTP Headers (request & response headers)

---

## 🖥️ 2. Go Back-end — πακέτο `net/http`
- [ ] `http.ListenAndServe(addr, nil)` — ανοίγω server σε port
- [ ] `http.HandleFunc(pattern, handler)` — routing
- [ ] Signature handler: `func(w http.ResponseWriter, r *http.Request)`
- [ ] `r.Method` — διαβάζω αν είναι GET ή POST
- [ ] `r.FormValue("key")` — παίρνω τιμή από HTML form
- [ ] `w.WriteHeader(statusCode)` — ορίζω status code
- [ ] `w.Write([]byte)` — γράφω body
- [ ] `http.Error(w, message, code)` — shorthand για error response
- [ ] `w.Header().Set("Key", "Value")` — ορίζω response header
- [ ] `http.Handle("/static/", http.FileServer(...))` — serve στατικά αρχεία (CSS)

---

## 🧩 3. Go Templates — πακέτο `html/template`
- [ ] `template.ParseFiles("templates/index.html")` — φορτώνω template
- [ ] `tmpl.Execute(w, data)` — εκτελώ template με δεδομένα
- [ ] Template syntax: `{{.FieldName}}` — τυπώνω field
- [ ] Template syntax: `{{if .Condition}} ... {{end}}` — conditional
- [ ] Template syntax: `{{range .Slice}} ... {{end}}` — loop
- [ ] Πώς περνώ struct ή map από Go στο HTML template

---

## 🌍 4. HTML Βασικά
- [ ] HTML skeleton: `<!DOCTYPE html>`, `<html>`, `<head>`, `<body>`
- [ ] `<form method="POST" action="/ascii-art">` — φτιάχνω HTML form
- [ ] `<input type="text">`, `<textarea>` — text inputs
- [ ] `<input type="radio">`, `<select>` — banner επιλογή
- [ ] `<button type="submit">` — submit form
- [ ] **`<pre>`** tag — monospace output για ASCII art (κρίσιμο!)
- [ ] `<link rel="stylesheet" href="/static/style.css">` — linking CSS

---

## 🎨 5. CSS Βασικά
- [ ] Selectors: element (`p`), class (`.btn`), id (`#output`)
- [ ] Box model: `margin`, `padding`, `border`, `width`, `height`
- [ ] Flexbox: `display: flex`, `flex-direction`, `align-items`, `justify-content`, `gap`
- [ ] `@media` queries για responsive design
- [ ] Transitions + `:hover` για interactive elements
- [ ] `font-family: monospace` για σωστή εμφάνιση ASCII art
- [ ] CSS variables με `--variable-name` (extra)

---

## 🐳 6. Docker
- [ ] Τι είναι Docker, image, container (η διαφορά)
- [ ] Γράψιμο `Dockerfile`: `FROM`, `WORKDIR`, `COPY`, `RUN`, `EXPOSE`, `CMD`
- [ ] `docker build -t name .` — χτίζω image
- [ ] `docker run -p 8080:8080 name` — τρέχω container
- [ ] `LABEL` instruction για metadata
- [ ] `.dockerignore` — τι αφήνω έξω από το image
- [ ] Multi-stage builds (μικρότερο τελικό image)
- [ ] Garbage collection: `docker container prune`, `docker image prune`, `docker system prune`

---

## 📤 7. HTTP Headers για Export
- [ ] `Content-Type: text/plain` — τι τύπος αρχείο στέλνουμε
- [ ] `Content-Disposition: attachment; filename="ascii-art.txt"` — download vs inline
- [ ] `Content-Length` — πόσα bytes έχει το αρχείο
- [ ] Πώς ορίζω headers σε Go **πριν** το `w.Write()`
- [ ] File permissions: `0644` (read+write user, read-only others)
- [ ] `/export` endpoint — νέο route που επιστρέφει αρχείο
