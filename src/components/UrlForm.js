// src/components/UrlForm.js
import React, { useState } from 'react';

function UrlForm({ onSubmit }) {
  const [originalUrl, setOriginalUrl] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(originalUrl);
    setOriginalUrl('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="url"
        placeholder="Enter URL to shorten"
        value={originalUrl}
        onChange={(e) => setOriginalUrl(e.target.value)}
        required
      />
      <button type="submit">Shorten URL</button>
    </form>
  );
}

export default UrlForm;
