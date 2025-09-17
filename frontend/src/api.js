import axios from "axios";

// User service
export const userService = axios.create({
  baseURL: "/api",
  headers: { "Content-Type": "application/json" },
});

export const setAuthToken = (token) => {
  if (token) {
    userService.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  } else {
    delete userService.defaults.headers.common["Authorization"];
  }
};

// Reimbursement service
export const reimbursementService = axios.create({
  baseURL: "/api/v1/reimbursements", // sesuaikan endpoint backend reimbursement
  headers: { "Content-Type": "application/json" },
});
