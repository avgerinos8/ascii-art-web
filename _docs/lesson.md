# 📘 Ολοκληρωμένο Μάθημα: HTML5 & CSS3 Essentials

---

## 🏗️ Ενότητα 1: Δομή & Σημαντική της HTML (Semantic HTML)

Η Semantic HTML χρησιμοποιεί ετικέτες που περιγράφουν το **περιεχόμενό** τους, βοηθώντας τις μηχανές αναζήτησης (SEO) και τους αναγνώστες οθόνης (Accessibility).

### 📝 Παράδειγμα Δομής Ιστοσελίδας

```html
<!DOCTYPE html>
<html lang="el">
<head>
    <meta charset="UTF-8">
    <title>Η Πρώτη Μου Σελίδα</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>

    <!-- Header: Κεφαλίδα με λογότυπο και πλοήγηση -->
    <header>
        <h1>Το Ιστολόγιό Μου</h1>
        <nav>
            <ul>
                <li><a href="#home">Αρχική</a></li>
                <li><a href="#blog">Άρθρα</a></li>
                <li><a href="#contact">Επικοινωνία</a></li>
            </ul>
        </nav>
    </header>

    <!-- Main: Το κύριο και μοναδικό περιεχόμενο της σελίδας -->
    <main>
        <!-- Section: Θεματική ενότητα -->
        <section id="blog">
            <!-- Article: Αυτόνομο περιεχόμενο (π.χ. ένα άρθρο) -->
            <article>
                <h2>Εισαγωγή στην HTML & CSS</h2>
                <p>Μαθαίνουμε πώς να χτίζουμε σύγχρονες ιστοσελίδες.</p>
            </article>
        </section>

        <!-- Aside: Πλευρική μπάρα με δευτερεύουσες πληροφορίες -->
        <aside>
            <h3>Σχετικά με μένα</h3>
            <p>Είμαι ένας Front-End Developer σε εξέλιξη!</p>
        </aside>
    </main>

    <!-- Footer: Υποσέλιδο με πνευματικά δικαιώματα -->
    <footer>
        <p>&copy; 2026 Το Όνομά Μου. All rights reserved.</p>
    </footer>

</body>
</html>
```

---

## 📝 Ενότητα 2: Φόρμες & Πεδία Εισαγωγής (Forms & Inputs)

Οι φόρμες επιτρέπουν τη συλλογή δεδομένων από τον χρήστη.

### 📝 Παράδειγμα Φόρμας Εγγραφής

```html
<form action="/submit-form" method="POST">
    
    <!-- Text Input & Label -->
    <div class="form-group">
        <label for="username">Όνομα Χρήστη:</label>
        <input type="text" id="username" name="username" placeholder="π.χ. Giannis99" required>
    </div>

    <!-- Email & Password Inputs -->
    <div class="form-group">
        <label for="user-email">Email:</label>
        <input type="email" id="user-email" name="email" required>
    </div>

    <div class="form-group">
        <label for="password">Κωδικός:</label>
        <input type="password" id="password" name="password" required>
    </div>

    <!-- Textarea για μεγάλα κείμενα -->
    <div class="form-group">
        <label for="bio">Βιογραφικό:</label>
        <textarea id="bio" name="bio" rows="4" placeholder="Λίγα λόγια για εσάς..."></textarea>
    </div>

    <!-- Dropdown Menu (Select & Option) -->
    <div class="form-group">
        <label for="country">Χώρα:</label>
        <select id="country" name="country">
            <option value="gr">Ελλάδα</option>
            <option value="cy">Κύπρος</option>
            <option value="other">Άλλη</option>
        </select>
    </div>

    <!-- Radio Buttons (Μονή επιλογή - ίδιο name) -->
    <div class="form-group">
        <p>Φύλο:</p>
        <input type="radio" id="male" name="gender" value="male">
        <label for="male">Άνδρας</label>
        
        <input type="radio" id="female" name="gender" value="female">
        <label for="female">Γυναίκα</label>
    </div>

    <!-- Checkboxes (Πολλαπλή επιλογή) -->
    <div class="form-group">
        <input type="checkbox" id="terms" name="terms" required>
        <label for="terms">Αποδέχομαι τους όρους χρήσης</label>
    </div>

    <!-- Buttons -->
    <button type="submit">Εγγραφή</button>

</form>
```

---

## 📊 Ενότητα 3: Πίνακες & Πολυμέσα (Tables & Media)

### 📝 Πίνακας Δεδομένων (Table)

```html
<table>
    <thead>
        <tr>
            <th>Μάθημα</th>
            <th>Ώρες</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>HTML5</td>
            <td>10 ώρες</td>
        </tr>
        <tr>
            <td>CSS3</td>
            <td>15 ώρες</td>
        </tr>
    </tbody>
</table>
```

### 📝 Ήχος, Βίντεο & Iframe

```html
<!-- Audio Player -->
<audio src="podcast.mp3" controls></audio>

<!-- Video Player -->
<video src="tutorial.mp4" controls width="400"></video>

<!-- Embedded Content (π.χ. Google Maps ή YouTube) -->
<iframe src="https://wikipedia.org" width="100%" height="300"></iframe>
```

---

## 🎨 Ενότητα 4: Βασικά Στοιχεία CSS (Styling & Layout)

Η CSS δίνει στυλ και καθορίζει τη διάταξη των στοιχείων της HTML.

### 📝 CSS Selectors, Box Model & Typography

```css
/* 1. Selectors (Επιλογείς) */
p { color: #333333; }                 /* Tag Selector */
.form-group { margin-bottom: 15px; } /* Class Selector (χρήση τελείας) */
#username { border: 1px solid red; }  /* ID Selector (χρήση δίεσης) */

/* 2. Box Model (Μοντέλο Κουτιού) */
.box {
    width: 300px;
    height: 150px;
    padding: 20px;          /* Εσωτερικό περιθώριο */
    border: 2px solid #000; /* Περίγραμμα */
    margin: 10px;           /* Εξωτερικό περιθώριο */
}

/* 3. Typography & Colors */
body {
    font-family: 'Arial', sans-serif;
    line-height: 1.6;
    background-color: #f4f4f4;
}
```

### 📝 Layout Systems: Flexbox & Grid

#### Flexbox (Για μονοδιάστατες διατάξεις - γραμμή ή στήλη)
```css
.nav-list {
    display: flex;
    justify-content: space-between; /* Μοιράζει το χώρο ενδιάμεσα */
    align-items: center;            /* Κεντράρει κάθετα */
    list-style: none;
}
```

#### CSS Grid (Για δισδιάστατες διατάξεις - γραμμές ΚΑΙ στήλες)
```css
.gallery {
    display: grid;
    grid-template-columns: repeat(3, 1fr); /* 3 ισόποσες στήλες */
    gap: 15px;                             /* Κενό ανάμεσα στα κουτιά */
}
```

### 📝 Interactions (Hover & Transitions)

Δημιουργούν ομαλά εφέ όταν ο χρήστης αλληλεπιδρά με τη σελίδα.

```css
.btn-submit {
    background-color: blue;
    color: white;
    padding: 10px 20px;
    border: none;
    cursor: pointer;
    /* Ομαλή μετάβαση για την αλλαγή χρώματος σε 0.3 δευτερόλεπτα */
    transition: background-color 0.3s ease; 
}

/* Hover Effect: Τι συμβαίνει όταν ακουμπάει το ποντίκι */
.btn-submit:hover {
    background-color: darkblue;
}
```


# 📘 Εξειδικευμένο Μάθημα: HTML5 Forms & Inputs

Οι φόρμες (`<form>`) είναι το βασικό μέσο αλληλεπίδρασης ενός χρήστη με μια ιστοσελίδα, καθώς επιτρέπουν τη συλλογή και την αποστολή δεδομένων.

---

## 1. Δομή & Στοιχεία της Φόρμας

### 🏗️ Το Κοντέινερ: `<form>`
Κλείνει μέσα του όλα τα πεδία εισαγωγής. Καθορίζει πού και πώς θα σταλούν τα δεδομένα.
*   **`action`**: Η διεύθυνση (URL) του διακομιστή (server) που θα δεχτεί τα δεδομένα.
*   **`method`**: Ο τρόπος αποστολής. Συνήθως `GET` (για αναζητήσεις) ή `POST` (για ασφαλή αποστολή/εγγραφή).

### 🏷️ Ετικέτες Πεδίων: `<label>`
Συνδέουν ένα κείμενο-οδηγό με το αντίστοιχο πεδίο. Βοηθούν στην προσβασιμότητα και αν κάνεις κλικ πάνω τους, ενεργοποιείται το πεδίο.
*   Η σύνδεση γίνεται ταυτίζοντας το **`for`** του label με το **`id`** του input.

---

## 2. Αναλυτικός Οδηγός Πεδίων (Inputs)

### 🔤 Πεδία Κειμένου (Text, Email, Password)
Χρησιμοποιούν το στοιχείο `<input>` αλλάζοντας την ιδιότητα `type`.
*   **`type="text"`**: Για απλό κείμενο (π.χ. Όνομα).
*   **`type="email"`**: Ελέγχει αυτόματα αν η τιμή έχει τη μορφή `name@domain.com`.
*   **`type="password"`**: Κρύβει τους χαρακτήρες με βούλες για ασφάλεια.
*   **`placeholder`**: Το αχνό κείμενο βοήθειας που εξαφανίζεται όταν γράφεις.
*   **`required`**: Κάνει το πεδίο υποχρεωτικό για να υποβληθεί η φόρμα.

### 📝 Μεγάλα Κείμενα: `<textarea>`
Αυτόνομο στοιχείο για κείμενα πολλών γραμμών (π.χ. Σχόλια, Μηνύματα).
*   Δεν χρησιμοποιεί `value` attribute. Το αρχικό κείμενο μπαίνει ανάμεσα στο άνοιγμα και το κλείσιμο: `<textarea>Κείμενο</textarea>`.
*   **`rows`** & **`cols`**: Καθορίζουν το αρχικό μέγεθος σε γραμμές και στήλες.

### 🔘 Radio Buttons & 🗹 Checkboxes
*   **`type="checkbox"`**: Για επιλογές "Ναι/Όχι" ή πολλαπλές απαντήσεις (μπορείς να διαλέξεις όσα θες).
*   **`type="radio"`**: Για αποκλειστικές επιλογές (διαλέγεις μόνο ΕΝΑ). Για να λειτουργήσει η αποκλειστικότητα, όλα τα radio buttons της ίδιας ομάδας πρέπει να έχουν το **ίδιο `name`**.

### 🔽 Λίστες Επιλογών: `<select>` & `<option>`
Δημιουργούν ένα dropdown μενού για εξοικονόμηση χώρου.
*   Το `<select>` είναι το περίβλημα και παίρνει το `name`.
*   Κάθε `<option>` είναι μια επιλογή και παίρνει μια κρυφή τιμή `value` που στέλνεται στον server.

### 🚀 Κουμπιά Υποβολής: `<button>` & `<input type="submit">`
*   **`<input type="submit" value="Αποστολή">`**: Παλιότερος τρόπος, δημιουργεί κουμπί με κείμενο αυτό που γράφεις στο `value`.
*   **`<button type="submit">Αποστολή</button>`**: Σύγχρονος και προτιμώμενος τρόπος. Μπορεί να κλείσει μέσα του κείμενο, εικόνες ή εικονίδια.

---

## 💻 Ολοκληρωμένο Παράδειγμα Κώδικα

```html
<form action="/submit-profile" method="POST">
    
    <!-- 1. Text & Label -->
    <div class="form-group">
        <label for="username">Όνομα Χρήστη:</label>
        <input type="text" id="username" name="username" placeholder="π.χ. thodoras99" required>
    </div>

    <!-- 2. Email & Password -->
    <div class="form-group">
        <label for="user-email">Email:</label>
        <input type="email" id="user-email" name="email" required>
    </div>

    <div class="form-group">
        <label for="user-pass">Κωδικός:</label>
        <input type="password" id="user-pass" name="password" required>
    </div>

    <!-- 3. Textarea -->
    <div class="form-group">
        <label for="user-bio">Λίγα λόγια για εσένα:</label>
        <textarea id="user-bio" name="bio" rows="4" cols="50"></textarea>
    </div>

    <!-- 4. Dropdown Menu -->
    <div class="form-group">
        <label for="user-role">Ρόλος:</label>
        <select id="user-role" name="role">
            <option value="student">Μαθητής</option>
            <option value="developer">Προγραμματιστής</option>
            <option value="hobbyist">Χομπίστας</option>
        </select>
    </div>

    <!-- 5. Radio Buttons (Μονή επιλογή - Ίδιο name) -->
    <div class="form-group">
        <p>Επίπεδο εμπειρίας:</p>
        
        <input type="radio" id="exp-beg" name="experience" value="beginner" checked>
        <label for="exp-beg">Αρχάριος</label>
        
        <input type="radio" id="exp-adv" name="experience" value="advanced">
        <label for="exp-adv">Προχωρημένος</label>
    </div>

    <!-- 6. Checkboxes (Πολλαπλή επιλογή) -->
    <div class="form-group">
        <input type="checkbox" id="newsletter" name="subscribe" value="yes">
        <label for="newsletter">Θέλω να λαμβάνω ενημερώσεις</label>
    </div>

    <!-- 7. Buttons -->
    <div class="form-actions">
        <!-- Κλασικό input submit -->
        <input type="submit" value="Γρήγορη Υποβολή">
        
        <!-- Σύγχρονο κουμπί button -->
        <button type="submit">Ολοκλήρωση Εγγραφής 🚀</button>
    </div>

</form>
```
# 📘 Εξειδικευμένο Μάθημα: Semantic HTML & Layout Structure

Η **Σημασιολογική HTML (Semantic HTML)** χρησιμοποιεί ετικέτες που περιγράφουν ξεκάθαρα το νόημα και τον ρόλο του περιεχομένου τους, τόσο στον προγραμματιστή όσο και στον browser ή στις μηχανές αναζήτησης (SEO).

Πριν την HTML5, χρησιμοποιούσαμε παντού `<div>` (π.χ. `<div class="header">`). Τώρα χρησιμοποιούμε ειδικές ετικέτες που βοηθούν τα άτομα με αναπηρία (μέσω screen readers) και βελτιώνουν την κατάταξη της σελίδας στο Google.

---

## 1. Ανάλυση Δομικών Στοιχείων (Layout Elements)

### 🔝 Η Κεφαλίδα: `<header>`
Αντιπροσωπεύει εισαγωγικό περιεχόμενο, συνήθως μια ομάδα στοιχείων πλοήγησης ή τίτλων.
*   **Περιέχει**: Το λογότυπο της ιστοσελίδας, τον κεντρικό τίτλο `<h1>`, τη μπάρα αναζήτησης ή το μενού.
*   *Σημείωση*: Μπορείς να έχεις παραπάνω από ένα `<header>` στη σελίδα (π.χ. ένα μέσα σε ένα άρθρο), αλλά το κυριότερο βρίσκεται στην κορυφή.

### 🗺️ Ο Πλοηγός: `<nav>`
Ορίζει ένα σύνολο συνδέσμων πλοήγησης (navigation links).
*   **Χρήση**: Το χρησιμοποιούμε για το κύριο μενού του site, για μενού στο υποσέλιδο ή για split-links (π.χ. προηγούμενη/επόμενη σελίδα).
*   **Tip**: Μέσα στο `<nav>` βάζουμε συνήθως μια μη ταξινομημένη λίστα `<ul>` με στοιχεία `<li>` και συνδέσμους `<a>`.

### 🎯 Το Κύριο Περιεχόμενο: `<main>`
Καθορίζει το μοναδικό, κεντρικό περιεχόμενο του εγγράφου.
*   **Κανόνας**: Επιτρέπεται **μόνο ένα** στοιχείο `<main>` ανά σελίδα.
*   **Περιεχόμενο**: Δεν πρέπει να περιλαμβάνει στοιχεία που επαναλαμβάνονται σε άλλες σελίδες (όπως sidebars, logos, copyrights).

### 📦 Θεματικές Ενότητες: `<section>`
Αντιπροσωπεύει ένα αυτόνομο τμήμα μιας σελίδας που έχει μια συγκεκριμένη θεματολογία.
*   **Χρήση**: Για να ομαδοποιήσεις σχετικό περιεχόμενο (π.χ. Ενότητα "Υπηρεσίες", Ενότητα "Επικοινωνία").
*   **Κανόνας**: Σχεδόν πάντα ένα `<section>` πρέπει να ξεκινάει με έναν τίτλο (`<h2>` έως `<h6>`).

### 📰 Αυτόνομα Άρθρα: `<article>`
Αντιπροσωπεύει μια αυτόνομη σύνθεση σε ένα έγγραφο, η οποία μπορεί να διαβαστεί και να σταθεί μόνη της ανεξάρτητα από το υπόλοιπο site.
*   **Χρήση**: Για posts σε blog, άρθρα ειδήσεων, forum posts, ή κάρτες προϊόντων (product cards).
*   **Tip**: Αν το περιεχόμενο μπορεί να γίνει "syndicated" (δηλαδή να το πάρει ένα άλλο site και να βγάζει νόημα αυτόνομο), τότε είναι `<article>`.

### 🧭 Η Πλευρική Μπάρα: `<aside>`
Ορίζει περιεχόμενο που σχετίζεται έμμεσα ή συμπληρωματικά με το κύριο περιεχόμενο γύρω του.
*   **Χρήση**: Για sidebar με links, διαφημίσεις, λίστες με "δημοφιλή άρθρα", βιογραφικό συγγραφέα ή ορισμούς (γλωσσάρι).

### 🔚 Το Υποσέλιδο: `<footer>`
Ορίζει το κάτω μέρος μιας σελίδας ή μιας ενότητας.
*   **Περιέχει**: Πληροφορίες πνευματικών δικαιωμάτων (copyright), συνδέσμους για όρους χρήσης, social media icons, ή στοιχεία επικοινωνίας.

---

## 💻 Ολοκληρωμένο Παράδειγμα Κώδικα Layout

Αυτό το παράδειγμα δείχνει πώς συνδυάζονται όλα τα παραπάνω στοιχεία για να σχηματίσουν τη σωστή "ραχοκοκαλιά" μιας σύγχρονης ιστοσελίδας:

```html
<!DOCTYPE html>
<html lang="el">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Η Σημασιολογική μου Ιστοσελίδα</title>
</head>
<body>

    <!-- 1. Κεφαλίδα Ιστοσελίδας -->
    <header>
        <div class="brand">TechBlog</div>
        
        <!-- 2. Κεντρικό Μενού Πλοήγησης -->
        <nav>
            <ul>
                <li><a href="#home">Αρχική</a></li>
                <li><a href="#news">Νέα</a></li>
                <li><a href="#contact">Επικοινωνία</a></li>
            </ul>
        </nav>
    </header>

    <!-- 3. Κύριο Περιεχόμενο (Μοναδικό) -->
    <main>

        <!-- 4. Θεματική Ενότητα για τα Άρθρα -->
        <section id="news">
            <h2>Τελευταίες Δημοσιεύσεις</h2>

            <!-- 5. Αυτόνομο Άρθρο 1 -->
            <article class="post">
                <h3>Τι είναι η Semantic HTML5;</h3>
                <p>Η Semantic HTML βοηθάει τις μηχανές αναζήτησης να καταλάβουν τη δομή του κώδικά μας...</p>
                <a href="/html-article">Διαβάστε περισσότερα</a>
            </article>

            <!-- Αυτόνομο Άρθρο 2 -->
            <article class="post">
                <h3>Οδηγός για το CSS Flexbox</h3>
                <p>Το Flexbox κάνει τη στοίχιση των στοιχείων παιχνιδάκι σε μία διάσταση...</p>
                <a href="/css-flexbox">Διαβάστε περισσότερα</a>
            </article>
        </section>

        <!-- 6. Πλευρική Μπάρα (Sidebar) -->
        <aside class="sidebar">
            <h3>Δημοφιλή Tags</h3>
            <ul>
                <li><a href="#">#html5</a></li>
                <li><a href="#">#css3</a></li>
                <li><a href="#">#webdev</a></li>
            </ul>
        </aside>

    </main>

    <!-- 7. Υποσέλιδο Ιστοσελίδας -->
    <footer>
        <p>&copy; 2026 TechBlog. Με επιφύλαξη παντός δικαιώματος.</p>
        <p><a href="/privacy">Πολιτική Απορρήτου</a></p>
    </footer>

</body>
</html>
```
# 📘 Εξειδικευμένο Μάθημα: Tables, Media & Embedded Content

Αυτή η ενότητα καλύπτει τη διαχείριση δομημένων δεδομένων μέσω πινάκων, την ενσωμάτωση αρχείων αναπαραγωγής ήχου και βίντεο, καθώς και την προβολή εξωτερικού περιεχομένου μέσα από τη σελίδα σας.

---

## 1. Πίνακες Δεδομένων (Tables)

Οι πίνακες χρησιμοποιούνται αποκλειστικά για την παρουσίαση **πίνακα δεδομένων** (tabular data) και όχι για τη σχεδίαση του layout της σελίδας. 

*   **`<table>`**: Το αρχικό κοντέινερ που ορίζει τον πίνακα.
*   **`<tr>` (Table Row)**: Ορίζει μια οριζόντια γραμμή στον πίνακα.
*   **`<th>` (Table Header)**: Ορίζει ένα κελί-επικεφαλίδα. Το κείμενο εμφανίζεται αυτόματα έντονο (bold) και κεντραρισμένο.
*   **`<td>` (Table Data)**: Ορίζει ένα κανονικό κελί δεδομένων.

### 💡 Προχωρημένη Σημαντική Δομή Πινάκων
Για μεγάλους πίνακες, χωρίζουμε το περιεχόμενο σε τρία μέρη για καλύτερη οργάνωση:
1.  **`<thead>`**: Περιέχει τις γραμμές με τις επικεφαλίδες.
2.  **`<tbody>`**: Περιέχει το κυρίως σώμα των δεδομένων.
3.  **`<code><tfoot>`**: Περιέχει γραμμές με αθροίσματα ή τελικά συμπεράσματα.

---

## 2. Πολυμέσα (Audio & Video)

Τα στοιχεία αυτά επιτρέπουν την αναπαραγωγή αρχείων χωρίς την ανάγκη εξωτερικών plugins.

### 🎵 Ήχος: `<audio>`
Επιτρέπει την ενσωμάτωση ηχητικού περιεχομένου (π.χ. `.mp3`, `.wav`).
*   **`controls`**: Ιδιότητα χωρίς τιμή. Αν δεν προστεθεί, ο player θα είναι κρυφός. Εμφανίζει τα κουμπιά Play, Pause και την ένταση.
*   **`<source>`**: Φωλιάζει μέσα στο audio. Καθορίζει το μονοπάτι του αρχείου (`src`) και τον τύπο του (`type="audio/mpeg"`).

### 🎬 Βίντεο: `<video>`
Επιτρέπει την προβολή αρχείων βίντεο (π.χ. `.mp4`, `.webm`).
*   **`controls`**: Εμφανίζει τη μπάρα ελέγχου του βίντεο.
*   **`width` / `height`**: Καθορίζει τις διαστάσεις του player.
*   **`poster`**: Δέχεται το URL μιας εικόνας που θα φαίνεται πριν ο χρήστης πατήσει το Play.

---

## 3. Ενσωματωμένο Περιεχόμενο: `<iframe>`

Το **Inline Frame (`<iframe>`)** χρησιμοποιείται για την εμφάνιση μιας άλλης ιστοσελίδας ή εξωτερικού περιεχομένου μέσα στην τρέχουσα σελίδα σας.

*   **`src`**: Η διεύθυνση URL της σελίδας που θέλετε να ενσωματώσετε.
*   **`title`**: Απαραίτητο για την προσβασιμότητα. Περιγράφει τι περιέχει το iframe.
*   **`loading="lazy"`**: Βελτιώνει την ταχύτητα της σελίδας, καθώς το iframe φορτώνει μόνο όταν ο χρήστης σκρολάρει κοντά του.
*   *Συχνή Χρήση*: Ενσωμάτωση χαρτών Google Maps, βίντεο από το YouTube ή widgets μέσων κοινωνικής δικτύωσης.

---

## 💻 Ολοκληρωμένο Παράδειγμα Κώδικα

```html
<!DOCTYPE html>
<html lang="el">
<head>
    <meta charset="UTF-8">
    <title>Πίνακες και Πολυμέσα στην HTML</title>
</head>
<body>

    <!-- 1. Παράδειγμα Πίνακα -->
    <section>
        <h2>Πρόγραμμα Μαθημάτων & Κόστος</h2>
        <table border="1"> <!-- Το border="1" βοηθάει να βλέπουμε τις γραμμές χωρίς CSS -->
            <thead>
                <tr>
                    <th>Μάθημα</th>
                    <th>Διάρκεια</th>
                    <th>Τιμή</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>Εισαγωγή στην HTML5</td>
                    <td>2 Εβδομάδες</td>
                    <td>50€</td>
                </tr>
                <tr>
                    <td>Προχωρημένη CSS3</td>
                    <td>3 Εβδομάδες</td>
                    <td>80€</td>
                </tr>
            </tbody>
            <tfoot>
                <tr>
                    <td colspan="2"><strong>Σύνολο:</strong></td>
                    <td><strong>130€</strong></td>
                </tr>
            </tfoot>
        </table>
    </section>

    <hr>

    <!-- 2. Παράδειγμα Ήχου -->
    <section>
        <h2>Ακούστε το τελευταίο μας Podcast</h2>
        <audio controls>
            <source src="podcast-intro.mp3" type="audio/mpeg">
            <source src="podcast-intro.ogg" type="audio/ogg">
            Το πρόγραμμα περιήγησής σας δεν υποστηρίζει το στοιχείο ήχου.
        </audio>
    </section>

    <hr>

    <!-- 3. Παράδειγμα Βίντεο -->
    <section>
        <h2>Βίντεο Παρουσίασης</h2>
        <video width="480" height="320" controls poster="thumbnail.jpg">
            <source src="welcome-tutorial.mp4" type="video/mp4">
            Το πρόγραμμα περιήγησής σας δεν υποστηρίζει το στοιχείο βίντεο.
        </video>
    </section>

    <hr>

    <!-- 4. Παράδειγμα Iframe -->
    <section>
        <h2>Η Τοποθεσία μας (Χάρτης)</h2>
        <iframe 
            src="https://wikipedia.org" 
            width="100%" 
            height="300" 
            title="Εγκυκλοπαίδεια Wikipedia" 
            loading="lazy">
            Το πρόγραμμα περιήγησής σας δεν υποστηρίζει iframes.
        </iframe>
    </section>

</body>
</html>
```
# 📘 Εξειδικευμένο Μάθημα: CSS3 Essentials for Layout & Styling

Η CSS (Cascading Style Sheets) είναι η γλώσσα που καθορίζει την εμφάνιση, τα χρώματα, τη γραμματοσειρά και τη διάταξη (layout) των στοιχείων της HTML. 

---

## 1. Επιλογείς (Selectors)

Οι Selectors είναι ο τρόπος με τον οποίο "στοχεύουμε" ένα HTML στοιχείο για να του δώσουμε στυλ.

*   **Tag Selector**: Στοχεύει όλα τα στοιχεία ενός συγκεκριμένου τύπου.
    ```css
    p { color: blue; } /* Όλες οι παράγραφοι γίνονται μπλε */
    ```
*   **Class Selector**: Στοχεύει στοιχεία που έχουν ένα συγκεκριμένο attribute `class=""`. Χρησιμοποιούμε μια **τελεία (`.`)** στην CSS. Μπορεί να επαναληφθεί σε πολλά στοιχεία.
    ```css
    .btn-primary { background-color: green; }
    ```
*   **ID Selector**: Στοχεύει ένα και μοναδικό στοιχείο που έχει το συγκεκριμένο attribute `id=""`. Χρησιμοποιούμε μια **δίεση (`#`)**.
    ```css
    #main-header { font-size: 24px; }
    ```

---

## 2. Το Μοντέλο Κουτιού (Box Model)

Κάθε στοιχείο στην HTML θεωρείται ένα ορθογώνιο κουτί. Το Box Model αποτελείται από 4 στρώματα, από μέσα προς τα έξω:

1.  **Content (Width & Height)**: Το ίδιο το περιεχόμενο (κείμενο ή εικόνα).
2.  **Padding**: Το εσωτερικό περιθώριο (ο χώρος ανάμεσα στο περιεχόμενο και το περίγραμμα).
3.  **Border**: Το περίγραμμα γύρω από το padding.
4.  **Margin**: Το εξωτερικό περιθώριο (ο χώρος που χωρίζει αυτό το κουτί από τα διπλανά του κουτιά).

*Σημαντικό Tip*: Χρησιμοποιούμε πάντα `box-sizing: border-box;` ώστε το padding και το border να μην μεγαλώνουν τις τελικές διαστάσεις του width.

---

## 3. Χρώματα & Τυπογραφία (Colors & Typography)

*   **`color`**: Αλλάζει το χρώμα του κειμένου (δέχεται ονόματα, HEX codes π.χ. `#ffffff`, ή RGB/RGBA).
*   **`background-color`**: Αλλάζει το χρώμα φόντου του κουτιού.
*   **`font-family`**: Καθορίζει τη γραμματοσειρά (π.χ. `Arial, sans-serif`).
*   **`font-size`**: Το μέγεθος των γραμμάτων (σε `px`, `em`, ή `rem`).

---

## 4. Συστήματα Διάταξης (Layout Systems)

### ↔️ Flexbox (Μονοδιάστατο Layout)
Ιδανικό για να στοιχίζεις στοιχεία σε μια σειρά (row) ή μια στήλη (column).
*   `display: flex;`: Μετρέπει το κοντέινερ σε flex container.
*   `justify-content`: Στοιχίζει τα στοιχεία στον οριζόντιο άξονα (π.χ. `center`, `space-between`).
*   `align-items`: Στοιχίζει τα στοιχεία στον κάθετο άξονα (π.χ. `center`).

### 🕸️ CSS Grid (Δισδιάστατο Layout)
Ιδανικό για πλήρη πλέγματα με γραμμές ΚΑΙ στήλες ταυτόχρονα.
*   `display: grid;`: Ενεργοποιεί το grid.
*   `grid-template-columns`: Ορίζει πόσες στήλες θα έχει το πλέγμα και τι μέγεθος (π.χ. `repeat(3, 1fr)` για 3 ίσες στήλες).
*   `gap`: Η απόσταση ανάμεσα στα κελιά του πλέγματος.

---

## 5. Αλληλεπιδράσεις (Interactions & Transitions)

*   **`:hover` (Pseudo-class)**: Εφαρμόζει στυλ μόνο όταν ο χρήστης περνάει το ποντίκι πάνω από το στοιχείο.
*   **`transition`**: Δημιουργεί ομαλή κίνηση (animation) κατά την αλλαγή μιας ιδιότητας (π.χ. από ένα χρώμα σε ένα άλλο), αντί η αλλαγή να γίνει ακαριαία.

---

## 💻 Ολοκληρωμένο Παράδειγμα Κώδικα

Ακολουθεί ο κώδικας HTML και CSS που δείχνει πώς εφαρμόζονται όλες οι παραπάνω έννοιες μαζί:

### HTML
```html
<!DOCTYPE html>
<html lang="el">
<head>
    <meta charset="UTF-8">
    <title>CSS Essentials Demo</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>

    <!-- ID Selector Example -->
    <header id="main-header">
        <h1>Καλώς ήρθατε στην CSS</h1>
    </header>

    <!-- Flexbox Container (Για το μενού) -->
    <nav class="flex-menu">
        <a href="#">Αρχική</a>
        <a href="#">Υπηρεσίες</a>
        <a href="#">Επικοινωνία</a>
    </nav>

    <!-- CSS Grid Container (Για τις κάρτες) -->
    <main class="grid-layout">
        
        <!-- Κάρτα 1 (Box Model & Interaction) -->
        <article class="card">
            <h3>Κάρτα HTML</h3>
            <p>Αυτό είναι ένα παράδειγμα Box Model.</p>
            <button class="btn-action">Δες Περισσότερα</button>
        </article>

        <!-- Κάρτα 2 -->
        <article class="card">
            <h3>Κάρτα CSS</h3>
            <p>Το background και τα borders αλλάζουν εδώ.</p>
            <button class="btn-action">Δες Περισσότερα</button>
        </article>

        <!-- Κάρτα 3 -->
        <article class="card">
            <h3>Κάρτα Grid</h3>
            <p>Αυτές οι 3 κάρτες μπήκαν δίπλα-δίπλα με Grid.</p>
            <button class="btn-action">Δες Περισσότερα</button>
        </article>

    </main>

</body>
</html>
```

### CSS (`style.css`)
```css
/* 1. Global Reset & Box Sizing */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box; 
}

/* 2. Typography & Colors (Tag Selector) */
body {
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    background-color: #f0f2f5;
    color: #333;
    padding: 20px;
}

/* ID Selector */
#main-header {
    text-align: center;
    margin-bottom: 20px;
}

/* 3. Flexbox Layout System */
.flex-menu {
    display: flex;
    justify-content: center;
    gap: 15px;
    background-color: #2c3e50;
    padding: 10px;
    border-radius: 8px;
}

.flex-menu a {
    color: white;
    text-decoration: none;
    padding: 5px 10px;
}

/* 4. CSS Grid Layout System */
.grid-layout {
    display: grid;
    grid-template-columns: repeat(3, 1fr); /* 3 ίσες στήλες */
    gap: 20px;
    margin-top: 30px;
}

/* 5. Box Model στην Πράξη (Class Selector) */
.card {
    background-color: white;
    padding: 20px;          /* Εσωτερικό περιθώριο */
    border: 1px solid #ddd; /* Περίγραμμα */
    border-radius: 6px;     /* Στρογγυλεμένες γωνίες */
    box-shadow: 0 4px 6px rgba(0,0,0,0.05);
}

.card h3 {
    margin-bottom: 10px; /* Απόσταση κάτω από τον τίτλο */
    color: #2c3e50;
}

/* 6. Interactions (Hover & Transitions) */
.btn-action {
    background-color: #3498db;
    color: white;
    border: none;
    padding: 10px 15px;
    margin-top: 15px;
    border-radius: 4px;
    cursor: pointer;
    
    /* Ομαλή αλλαγή για το background και το transform μέσα σε 0.3 δευτερόλεπτα */
    transition: background-color 0.3s ease, transform 0.2s ease;
}

/* Hover Effect */
.btn-action:hover {
    background-color: #2980b9; /* Σκούρο μπλε στο hover */
    transform: translateY(-2px); /* Μικρή ανύψωση του κουμπιού */
}
```
