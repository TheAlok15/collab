import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import api from "../api/Axios.jsx";


export default function SignupPage() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await api.post("/signup",{
        name,
        email,
        password
      })

      alert("Signup successful! Please sign in.");
      navigate("/signin")
      
    } catch (error) {
       if (error.response) {
        alert(error.response.data.error || "Signup failed");
      } else {
        alert("Network error");
      }
      
    }
  };

  return (
    <div className="min-h-screen flex justify-center items-center p-4 bg-gray-100">
      <form
        onSubmit={handleSubmit}
        className="w-full max-w-md bg-white shadow-xl rounded-2xl p-6 flex flex-col gap-4"
      >
        <h2 className="text-2xl font-semibold text-gray-900 text-center">
          Create Account
        </h2>

        {/* Name */}
        <div className="flex flex-col w-full gap-1">
          <label className="text-sm font-medium text-gray-700">Full Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="John Doe"
            className="w-full px-3 py-2 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all"
          />
        </div>

        {/* Email */}
        <div className="flex flex-col w-full gap-1">
          <label className="text-sm font-medium text-gray-700">Email</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="you@example.com"
            className="w-full px-3 py-2 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all"
          />
        </div>

        {/* Password */}
        <div className="flex flex-col w-full gap-1">
          <label className="text-sm font-medium text-gray-700">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="••••••••"
            className="w-full px-3 py-2 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all"
          />
        </div>

        <button
          type="submit"
          className="w-full py-2 bg-indigo-600 text-white rounded-xl font-medium hover:bg-indigo-700 active:scale-[0.98] transition-all"
        >
          Sign Up
        </button>
        <p className="text-sm text-center text-gray-600">
          Already signed up?{" "}
          <Link to="/signin" className="text-indigo-600 hover:underline">
            Sign in
          </Link>
        </p>
      </form>
    </div>
  );
}
