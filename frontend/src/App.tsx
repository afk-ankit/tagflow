import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Dashboard from './components/Dashboard';
import Deployments from './components/Deployments';
import './index.css';

const App: React.FC = () => {
  return (
    <Router>
      <div className="app-container">
        <nav>
          <ul>
            <li>
              <Link to="/">Dashboard</Link>
            </li>
            <li>
              <Link to="/deployments">Deployments</Link>
            </li>
          </ul>
        </nav>

        <div className="content">
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/deployments" element={<Deployments />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default App;
