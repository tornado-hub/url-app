import React, { useState } from 'react';

function UrlForm({ onSubmit }) {
  const [originalUrl, setOriginalUrl] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(originalUrl);
    setOriginalUrl('');
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4">
      <div className="input-group">
        <input
          type="url"
          className="form-control"
          placeholder="Enter URL to shorten"
          value={originalUrl}
          onChange={(e) => setOriginalUrl(e.target.value)}
          required
        />
        <div className="input-group-append">
          <button className="btn btn-success" type="submit">
            Shorten URL
          </button>
        </div>
      </div>
    </form>
  );
}

export default UrlForm;
