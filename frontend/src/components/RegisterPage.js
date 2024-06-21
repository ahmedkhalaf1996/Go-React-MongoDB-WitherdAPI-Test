import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { register } from '../api';

const RegisterPage = () => {
  const [credentials, setCredentials] = useState({
    username: '',
    email: '',
    password: '',
    lat: '',
    lon: '',
  });
  const navigate = useNavigate();

  useEffect(() => {
    navigator.geolocation.getCurrentPosition(
      (position) => {
        setCredentials((prevCreds) => ({
          ...prevCreds,
          lat: position.coords.latitude,
          lon: position.coords.longitude,
        }));
      },
      (error) => {
        console.error('Error getting geolocation:', error);
      }
    );
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCredentials((prevCreds) => ({ ...prevCreds, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await register(credentials);
      navigate('/login');
    } catch (err) {
      console.error(err);
      alert('Registration failed.');
    }
  };


  return (
    <div className="container">
      <h1>Register</h1>
      <form onSubmit={handleSubmit}>
        <input
          name="username"
          value={credentials.username}
          onChange={handleChange}
          placeholder="Username"
          required
        />
        <input
          name="email"
          type="email"
          value={credentials.email}
          onChange={handleChange}
          placeholder="Email"
          required
        />
        <input
          name="password"
          type="password"
          value={credentials.password}
          onChange={handleChange}
          placeholder="Password"
          required
        />
        <button type="submit">Register</button>
      </form>
      <p>
        Already have an account?{' '}
        <button className="link-button" onClick={() => navigate('/login')}>
          Login
        </button>
      </p>
    </div>
  );
};

export default RegisterPage;
