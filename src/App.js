import React, { useState, useEffect } from 'react';
import UrlForm from './components/UrlForm';
import UrlList from './components/UrlList';

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
        body: JSON.stringify({ original_url: originalUrl }), // Ensure 'original_url' matches backend JSON key
      });

      if (response.ok) {
        fetchUrls(); // Refresh the list of URLs
      } else {
        console.error('Error shortening URL:', response.statusText);
      }
    } catch (error) {
      console.error('Error submitting URL:', error);
    }
  };

  useEffect(() => {
    fetchUrls();
  }, []);

  return (
    <div>
      <h1>URL Shortener</h1>
      <UrlForm onSubmit={handleSubmit} />
      <UrlList urls={urls} />
    </div>
  );
}

export default App;
