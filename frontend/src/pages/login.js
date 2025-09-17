import React, { useState } from "react";
import { userService, setAuthToken } from "../api";
import "./login.css"; // Kita pakai CSS terpisah

function Login({ onLogin }) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await userService.post("/v1/auth/login", { email, password });
      const { token, user } = res.data;
      setAuthToken(token);
      onLogin(token, user);
    } catch (err) {
      alert("Login gagal!");
      console.error(err);
    }
  };

  return (
    <div className="login-container">
      <div className="login-card">
        <h2 className="login-title">Welcome Back</h2>
        <form onSubmit={handleSubmit}>
          <input
            type="email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            placeholder="Email"
            required
            className="login-input"
          />
          <input
            type="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            placeholder="Password"
            required
            className="login-input"
          />
          <button type="submit" className="login-button">Login</button>
        </form>
        <p className="login-footer">© 2025 Your Company</p>
      </div>
    </div>
  );
}

export default Login;
