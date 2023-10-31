package main

import (
	"fmt"
	"net/http"
)

// multiplexer
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html>

<head>
    <title>Home Page</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
        }

        .header {
            background-color: #333;
            color: #fff;
            text-align: center;
            padding: 30px 0;
        }

        .navbar {
            overflow: hidden;
            background-color: #333;
        }

        .navbar a {
            float: left;
            display: block;
            color: #f2f2f2;
            text-align: center;
            padding: 14px 20px;
            text-decoration: none;
        }

        .navbar a:hover {
            background-color: #ddd;
            color: black;
        }

        .row {
            display: flex;
        }

        .column {
            flex: 50%;
            padding: 10px;
        }

        .column img {
            width: 100%;
            height: auto;
        }

        .footer {
            background-color: #333;
            color: #fff;
            text-align: center;
            padding: 20px 0;
            position: absolute;
            bottom: 0;
            width: 100%;
        }
    </style>
</head>

<body>
    <div class="header">
        <h1>Our Website</h1>
    </div>

    <div class="navbar">
        <a href="#">Home</a>
        <a href="#">About</a>
        <a href="#">Services</a>
        <a href="#">Contact</a>
    </div>

    <div class="row">
        <div class="column">
            <img src="https://picsum.photos/200/300
" alt="Image 1">
        </div>
        <div class="column">
            <img src="https://picsum.photos/200/300
" alt="Image 2">
        </div>
    </div>

    <div class="footer">
        <p>&copy; 2023 Our Website</p>
    </div>
</body>

</html>
`)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html>

<head>
    <title>Welcome Page</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f2f2f2;
        }

        .welcome-box {
            text-align: center;
            padding: 50px;
            background-color: #ffffff;
            border-radius: 10px;
            box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.1);
        }

        h1 {
            color: #333;
            margin-bottom: 30px;
        }

        p {
            color: #666;
            font-size: 18px;
            line-height: 1.6;
        }

        .button {
            display: inline-block;
            padding: 10px 20px;
            margin-top: 30px;
            color: #fff;
            text-transform: uppercase;
            text-decoration: none;
            background-color: #ff7b29;
            border-radius: 5px;
            transition: 0.5s;
        }

        .button:hover {
            background-color: #e65c00;
        }
    </style>
</head>

<body>
    <div class="welcome-box">
        <h1>Welcome to Our Website</h1>
        <p>Thank you for visiting our page. We provide the best services and products to meet your needs. Feel free to explore and reach out to us for any inquiries.</p>
        <a class="button" href="#">Get Started</a>
    </div>
</body>

</html>
`)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html>

<head>
    <title>Login Form</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .login-box {
            width: 300px;
            padding: 40px;
            background: #fff;
            border-radius: 10px;
            box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.1);
            text-align: center;
        }

        .login-box h2 {
            margin-bottom: 30px;
            text-transform: uppercase;
            color: #333;
        }

        .user-box {
            position: relative;
            margin-bottom: 30px;
        }

        .user-box input {
            width: 100%;
            padding: 10px 0;
            font-size: 16px;
            color: #333;
            margin-bottom: 30px;
            border: none;
            border-bottom: 1px solid #999;
            outline: none;
            background: transparent;
        }

        .user-box label {
            position: absolute;
            top: 0;
            left: 0;
            padding: 10px 0;
            font-size: 16px;
            color: #999;
            pointer-events: none;
            transition: 0.5s;
        }

        .user-box input:focus~label,
        .user-box input:valid~label {
            top: -20px;
            left: 0;
            color: #ff7b29;
            font-size: 12px;
        }

        a {
            display: inline-block;
            padding: 10px 20px;
            color: #fff;
            text-transform: uppercase;
            text-decoration: none;
            background: #ff7b29;
            border-radius: 5px;
            transition: 0.5s;
        }

        a:hover {
            background: #333;
        }
    </style>
</head>

<body>
    <div class="login-box">
        <h2>Login</h2>
        <form>
            <div class="user-box">
                <input type="text" name="" required="">
                <label>Username</label>
            </div>
            <div class="user-box">
                <input type="password" name="" required="">
                <label>Password</label>
            </div>
            <a href="#">
                <span></span>
                <span></span>
                <span></span>
                <span></span>
                Submit
            </a>
        </form>
    </div>
</body>

</html>
`)
}

func main() {
	http.Handle("/home", http.HandlerFunc(home))
	http.Handle("/login", http.HandlerFunc(login))
	http.Handle("/welcome", http.HandlerFunc(welcome))
	http.ListenAndServe("localhost:8000", nil)

}
