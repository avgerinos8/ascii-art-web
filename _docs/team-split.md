# 👥 Team Split — pavgerin · cnewman · nantoniou
> 4 ασκήσεις, 3 άτομα. Κάθε άσκηση σε 3 ισόποσα κομμάτια.
> **Rotation principle**: κάθε άτομο αγγίζει back-end, front-end, **και** Docker.

---

## 📦 Άσκηση 1 — `ascii-art-web` (Web Server + HTML)

| Άτομο | Αρμοδιότητα | Αρχεία |
|---|---|---|
| **pavgerin** | Go server: `main.go`, routing, handler stubs, port setup, static file serving | `main.go` |
| **cnewman** | HTML template: form layout, text input, banner select/radio, `<pre>` output area | `templates/index.html` |
| **nantoniou** | ASCII art integration, status codes (200/400/404/500), error handling, `README.md` | `ascii.go`, `README.md` |

> Σημείωση: Ο cnewman χρειάζεται να συνεννοηθεί με τον pavgerin για τα ονόματα των form fields.

---

## 🐳 Άσκηση 2 — `ascii-art-web-dockerize`

| Άτομο | Αρμοδιότητα | Αρχεία |
|---|---|---|
| **cnewman** | Dockerfile: multi-stage build (builder + runner stage), `EXPOSE`, `CMD` | `Dockerfile` |
| **nantoniou** | `LABEL` metadata (authors, version, description), image size optimization | `Dockerfile` (labels section) |
| **pavgerin** | `.dockerignore`, build/run testing, garbage collection commands, Docker section στο README | `.dockerignore`, `README.md` update |

---

## 📤 Άσκηση 3 — `ascii-art-web-export`

| Άτομο | Αρμοδιότητα | Αρχεία |
|---|---|---|
| **nantoniou** | Go handler: `/export` endpoint, query param parsing, file generation | `handlers.go` (export func) |
| **pavgerin** | HTTP headers: `Content-Type`, `Content-Disposition`, `Content-Length` | `handlers.go` (header setup) |
| **cnewman** | HTML: download button/link, hidden form fields για το export, error feedback UI | `templates/index.html` update |

---

## 🎨 Άσκηση 4 — `ascii-art-web-stylize`

| Άτομο | Αρμοδιότητα | Αρχεία |
|---|---|---|
| **cnewman** | CSS base: typography, colors, monospace font για ASCII art, button styles | `static/style.css` (base) |
| **pavgerin** | Responsive design: `@media` queries, mobile layout, fluid widths | `static/style.css` (media queries) |
| **nantoniou** | Interactive elements: `:hover`, `:focus`, transitions, loading feedback, error states | `static/style.css` (interactive) |

---

## 📊 Rotation Summary

|  | Ex.1 | Ex.2 | Ex.3 | Ex.4 |
|---|---|---|---|---|
| **pavgerin** | 🖥️ Back-end / Routing | 🐳 Cleanup / .dockerignore | 📋 HTTP Headers | 📱 Responsive CSS |
| **cnewman** | 🌍 HTML Template | 🐳 Dockerfile | 🔽 Download UI | 🎨 CSS Base |
| **nantoniou** | ⚙️ Logic / Errors / README | 🐳 Labels / Metadata | 📤 Export Handler | ✨ Interactive CSS |

---

## 🔄 Σύσταση Workflow ανά Άσκηση

```
1. Ανοίξτε Discord call
2. Διαβάστε μαζί τις απαιτήσεις της άσκησης (5 λεπτά)
3. Ο καθένας ξεκινάει το κομμάτι του σε ξεχωριστό branch:
       git checkout -b ex1/pavgerin-server
       git checkout -b ex1/cnewman-template
       git checkout -b ex1/nantoniou-logic
4. Commit + push συχνά — μικρά commits με περιγραφικά μηνύματα
5. Όταν είστε έτοιμοι: code review μεταξύ σας, μετά merge στο main
```

> ⚠️ **Ποτέ** μη δουλεύετε απευθείας στο `main` branch.
