import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

export interface Project {
  id: number;
  name: string;
  branches: Branch[];
}

export interface Branch {
  id: number;
  name: string;
  type: string;
  tags: Tag[];
}

export interface Tag {
  id: number;
  name: string;
  created_by: string;
}

export interface Deployment {
  id: number;
  tag: Tag;
  environment: string;
  deployed_at: string;
}

export const getProjects = async () => {
  const response = await axios.get<Project[]>(`${API_URL}/projects`);
  return response.data;
};

export const getDeployments = async () => {
  const response = await axios.get<Deployment[]>(`${API_URL}/deployments`);
  return response.data;
};
