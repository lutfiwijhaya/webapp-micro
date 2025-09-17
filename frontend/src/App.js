import React, { useState } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/login";
import Reimbursements from "./pages/Reimbursements";
import SubmitReimbursement from "./pages/SubmitReimbursement";
import Navbar from "./components/Navbar";

export default function App() {
  const [user, setUser] = useState(null);

  const handleLogin = (token, userData) => setUser(userData);
  const handleLogout = () => setUser(null);

  if (!user) return <Login onLogin={handleLogin} />;

  return (
    <Router>
      <Navbar user={user} onLogout={handleLogout} />
      <Routes>
        <Route path="/" element={<Navigate to="/reimbursements" />} />
        <Route path="/reimbursements" element={<Reimbursements user={user} />} />
        <Route path="/submit" element={<SubmitReimbursement />} />
      </Routes>
    </Router>
  );
}
