import React, { useEffect, useState } from "react";
import { reimbursementService } from "../api";

export default function Reimbursements({ user }) {
  const [reimbursements, setReimbursements] = useState([]);

  const fetchData = async () => {
    try {
      const res = await reimbursementService.get("/api/v1/reimbursements");
      setReimbursements(res.data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleApprove = async (id) => {
    try {
      await reimbursementService.put(`/api/v1/reimbursements/approve/${id}`);
      fetchData();
    } catch (err) {
      console.error(err);
    }
  };

  const handleReject = async (id) => {
    try {
      await reimbursementService.put(`/api/v1/reimbursements/reject/${id}`);
      fetchData();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h1>Reimbursements</h1>
      <ul>
        {reimbursements.map(r => (
          <li key={r.id}>
            {r.title} - {r.status} - {r.amount}
            {(user.role === "manager" || user.role === "admin") && r.status === "pending" && (
              <>
                <button onClick={() => handleApprove(r.id)}>Approve</button>
                <button onClick={() => handleReject(r.id)}>Reject</button>
              </>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}
