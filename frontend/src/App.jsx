import { useState } from 'react'
import './App.css'
import SignupPage from './pages/Signup.jsx'
import SigninPage from './pages/Signin.jsx'
import { Route, Routes } from 'react-router-dom'
import Homepage from './pages/Homepage.jsx'


function App() {

  return (
    <Routes>
      <Route path="/" element={<SignupPage />} />
      <Route path="/signin" element={<SigninPage />} />
      <Route path="/home" element={<Homepage />} />
    </Routes>
  )
}

export default App
