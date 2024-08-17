//TODO: UI for login page, assuming SSH login can be input here, or perhaps a different method of authorization (admin login). 
import React from "react";
import Login from "@/components/ui/login"; // Ensure the path is correct

const LoginPage = () => {
  return (
    <div className="p-8">
      <h1 className="text-2xl font-bold mb-4">SSH Login Details</h1>
      <Login />
    </div>
  );
};

export default LoginPage;
