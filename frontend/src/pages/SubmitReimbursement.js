import React, { useState } from "react";
import { reimbursementService } from "../api";

export default function SubmitReimbursement() {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState("");
  const [category_id, setCategoryId] = useState("");
  const [file, setFile] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append("title", title);
    formData.append("description", description);
    formData.append("amount", amount);
    formData.append("category_id", category_id);
    formData.append("file", file);

    try {
      await reimbursementService.post("/api/v1/reimbursements/", formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      alert("Submitted!");
    } catch (err) {
      console.error(err);
      alert("Submit failed!");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" placeholder="Title" value={title} onChange={e => setTitle(e.target.value)} required />
      <textarea placeholder="Description" value={description} onChange={e => setDescription(e.target.value)} />
      <input type="number" placeholder="Amount" value={amount} onChange={e => setAmount(e.target.value)} required />
      <input type="number" placeholder="Category ID" value={category_id} onChange={e => setCategoryId(e.target.value)} required />
      <input type="file" onChange={e => setFile(e.target.files[0])} required />
      <button type="submit">Submit</button>
    </form>
  );
}
