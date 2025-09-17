import React from "react";

export default function Navbar({ user, onLogout }) {
  return (
    <nav>
      <span>Hello, {user.name} ({user.role})</span>
      <button onClick={onLogout}>Logout</button>
    </nav>
  );
}
