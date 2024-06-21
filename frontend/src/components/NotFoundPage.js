import React from 'react';
import { Link } from 'react-router-dom';

const NotFoundPage = () => {
  return (
    <div className="container">
      <h2>Page Not Found</h2>
      <p>
        Sorry, the page you are looking for does not exist. Go back to the{' '}
        <Link to="/login" className="link-button">
          Login Page
        </Link>
        .
      </p>
    </div>
  );
};

export default NotFoundPage;
