import React from 'react';

function UrlList({ urls }) {
  return (
    <ul>
      {urls.map((url) => (
        <li key={url.short_url}>
          <a href={url.original_url} target="_blank" rel="noopener noreferrer">
            {url.short_url}
          </a>
        </li>
      ))}
    </ul>
  );
}

export default UrlList;