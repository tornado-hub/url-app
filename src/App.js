import React, { useState, useEffect } from 'react';
import UrlForm from './components/UrlForm';
import UrlList from './components/UrlList';
import './index.css'; // Import the custom CSS

function App() {
  const [urls, setUrls] = useState([]);

  const fetchUrls = async () => {
    try {
      const response = await fetch('http://localhost:8000/urls');
      const data = await response.json();
      setUrls(data);
    } catch (error) {
      console.error('Error fetching URLs:', error);
    }
  };

  const handleSubmit = async (originalUrl) => {
    try {
      const response = await fetch('http://localhost:8000/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ original_url: originalUrl }),
      });

      if (response.ok) {
        fetchUrls();
      } else {
        console.error('Error shortening URL:', response.statusText);
      }
    } catch (error) {
      console.error('Error submitting URL:', error);
    }
  };

  const handleDelete = async (shortUrl) => {
    try {
      const response = await fetch('http://localhost:8000/delete', {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ short_url: shortUrl }),
      });

      if (response.ok) {
        fetchUrls();
      } else {
        console.error('Error deleting URL:', response.statusText);
      }
    } catch (error) {
      console.error('Error deleting URL:', error);
    }
  };

  useEffect(() => {
    fetchUrls();
  }, []);

  return (
    <div className="container">
      <h1 className="text-center my-4">URL Shortener</h1>
      <UrlForm onSubmit={handleSubmit} />
      <UrlList urls={urls} onDelete={handleDelete} />
    </div>
  );
}

export default App;
