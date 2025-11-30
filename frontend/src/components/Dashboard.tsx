import React, { useEffect, useState } from 'react';
import { getProjects, type Project } from '../api';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"

const Dashboard: React.FC = () => {
    const [projects, setProjects] = useState<Project[]>([]);

    useEffect(() => {
        getProjects().then(setProjects).catch(console.error);
    }, []);

    return (
        <div className="space-y-6">
            <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                {projects.map((project) => (
                    <Card key={project.id}>
                        <CardHeader>
                            <CardTitle>{project.name}</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <div className="space-y-4">
                                <h3 className="text-sm font-medium text-muted-foreground">Branches</h3>
                                <ul className="space-y-2">
                                    {project.branches?.map((branch) => (
                                        <li key={branch.id} className="flex flex-col space-y-1">
                                            <div className="flex items-center justify-between">
                                                <span className="font-medium">{branch.name}</span>
                                                <Badge variant="secondary">{branch.type}</Badge>
                                            </div>
                                            {branch.tags && branch.tags.length > 0 && (
                                                <div className="flex flex-wrap gap-1 mt-1">
                                                    {branch.tags.map(tag => (
                                                        <Badge key={tag.id} variant="outline" className="text-xs">{tag.name}</Badge>
                                                    ))}
                                                </div>
                                            )}
                                        </li>
                                    ))}
                                </ul>
                            </div>
                        </CardContent>
                    </Card>
                ))}
            </div>
        </div>
    );
};

export default Dashboard;
