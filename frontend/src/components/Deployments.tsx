import React, { useEffect, useState } from 'react';
import { getDeployments, type Deployment } from '../api';
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"

const Deployments: React.FC = () => {
    const [deployments, setDeployments] = useState<Deployment[]>([]);

    useEffect(() => {
        getDeployments().then(setDeployments).catch(console.error);
    }, []);

    return (
        <div className="space-y-6">
            <h1 className="text-3xl font-bold tracking-tight">Deployment History</h1>
            <div className="rounded-md border">
                <Table>
                    <TableCaption>A list of your recent deployments.</TableCaption>
                    <TableHeader>
                        <TableRow>
                            <TableHead>Tag</TableHead>
                            <TableHead>Environment</TableHead>
                            <TableHead className="text-right">Deployed At</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {deployments.map((d) => (
                            <TableRow key={d.id}>
                                <TableCell className="font-medium">{d.tag?.name}</TableCell>
                                <TableCell>{d.environment}</TableCell>
                                <TableCell className="text-right">{new Date(d.deployed_at).toLocaleString()}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>
        </div>
    );
};

export default Deployments;
