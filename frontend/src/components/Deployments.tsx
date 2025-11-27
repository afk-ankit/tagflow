import React, { useEffect, useState } from 'react';
import { getDeployments, type Deployment } from '../api';

const Deployments: React.FC = () => {
    const [deployments, setDeployments] = useState<Deployment[]>([]);

    useEffect(() => {
        getDeployments().then(setDeployments).catch(console.error);
    }, []);

    return (
        <div className="deployments">
            <h1>Deployment History</h1>
            <table>
                <thead>
                    <tr>
                        <th>Tag</th>
                        <th>Environment</th>
                        <th>Deployed At</th>
                    </tr>
                </thead>
                <tbody>
                    {deployments.map((d) => (
                        <tr key={d.id}>
                            <td>{d.tag?.name}</td>
                            <td>{d.environment}</td>
                            <td>{new Date(d.deployed_at).toLocaleString()}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default Deployments;
