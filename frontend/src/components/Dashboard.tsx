import React, { useEffect, useState } from 'react';
import { getProjects, type Project } from '../api';

const Dashboard: React.FC = () => {
    const [projects, setProjects] = useState<Project[]>([]);

    useEffect(() => {
        getProjects().then(setProjects).catch(console.error);
    }, []);

    return (
        <div className="dashboard">
            <h1>Tagflow Dashboard</h1>
            <div className="project-list">
                {projects.map((project) => (
                    <div key={project.id} className="project-card">
                        <h2>{project.name}</h2>
                        <div className="branches">
                            <h3>Branches</h3>
                            <ul>
                                {project.branches?.map((branch) => (
                                    <li key={branch.id}>
                                        {branch.name} ({branch.type})
                                        <ul>
                                            {branch.tags?.map(tag => (
                                                <li key={tag.id}>{tag.name}</li>
                                            ))}
                                        </ul>
                                    </li>
                                ))}
                            </ul>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Dashboard;
