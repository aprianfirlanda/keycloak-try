import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import keycloak from "./keycloak.js";

keycloak.init({ onLoad: "login-required" }).then((authenticated) => {
  if (authenticated) {
    createRoot(document.getElementById('root')).render(
      <StrictMode>
        <App keycloak={keycloak}/>
      </StrictMode>,
    )
  }
});
