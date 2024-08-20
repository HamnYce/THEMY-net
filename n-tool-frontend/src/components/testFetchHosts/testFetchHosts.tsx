import React, { useEffect, useState } from 'react';
import { fetchHosts } from '../../services/fetchData';

interface Host {
  id: number;
  name: string;
  ip: string;
  ports: string;
}

const TestFetchHosts = () => {
  const [hosts, setHosts] = useState<Host[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await fetchHosts(10, 0);
        if (data.errors && data.errors.length > 0) {
          setError(data.errors.join(', '));
        } else {
          setHosts(data.hosts);
        }
      } catch (err) {
        setError('Failed to fetch hosts');
      }
    };

    fetchData();
  }, []);

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      <h1>Hosts</h1>
      <ul>
        {hosts.map((host) => (
          <li key={host.id}>
            {host.name} - {host.ip} - {host.ports}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TestFetchHosts;
