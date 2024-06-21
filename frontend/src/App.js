import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import LoginPage from './components/LoginPage';
import RegisterPage from './components/RegisterPage';
import Dashboard from './components/Dashboard';
import NotFoundPage from './components/NotFoundPage';
import LogoutButton from './components/LogoutButton'; // Import the LogoutButton component

const App = () => {
  return (
    <Router>
      <div className="App">
        <header>
          <nav>
            <ul>
              <li><LogoutButton /></li>
            </ul>
          </nav>
        </header>
        <main>
          <Routes>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/" element={<Navigate to="/login" />} />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
};

export default App;
