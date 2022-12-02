import * as React from "react";
import { Routes, Route, Link } from "react-router-dom"
import Login from './pages/Login'
import Logout from './pages/Logout'
import ProtectedLayout from "./pages/layouts/ProtectedLayout"

export default function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Login />} />

        <Route path="/main" element={<ProtectedLayout />}>
          <Route index element={<Home />} />
          <Route path="about" element={<About />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="logout" element={<Logout />} />
          <Route path="*" element={<NoMatch />} />
        </Route>
      </Routes>
    </div>
  );
}

function Home() {
  return (
    <div>
      <h2>Home</h2>
    </div>
  );
}

function About() {
  return (
    <div>
      <h2>About</h2>
    </div>
  );
}

function Dashboard() {
  return (
    <div>
      <h2>Dashboard</h2>
    </div>
  );
}

function NoMatch() {
  return (
    <div>
      <h2>Nothing to see here!</h2>
      <p>
        <Link to="/">Go to the home page</Link>
      </p>
    </div>
  );
}