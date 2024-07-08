import React from 'react';

function UrlList({ urls, onDelete }) {
  const handleDelete = (shortUrl) => {
    onDelete(shortUrl);
  };

  return (
    <table className="table table-striped">
      <thead>
        <tr>
          <th>Short URL</th>
          <th>Original URL</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {urls.map((url) => (
          <tr key={url.short_url}>
            <td>
              <a href={`http://localhost:8000/${url.short_url}`} target="_blank" rel="noopener noreferrer">
                {`http://localhost:8000/${url.short_url}`}
              </a>
            </td>
            <td>
              <a href={url.original_url} target="_blank" rel="noopener noreferrer">
                {url.original_url}
              </a>
            </td>
            <td>
              <button className="btn btn-danger" onClick={() => handleDelete(url.short_url)}>
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default UrlList;
