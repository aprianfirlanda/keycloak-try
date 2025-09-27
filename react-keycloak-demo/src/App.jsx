function App({keycloak}) {
  return (
    <div>
      <h1>Welcome, {keycloak.tokenParsed?.preferred_username}</h1>
      <button onClick={() => keycloak.logout()}>Logout</button>
    </div>
  );
}

export default App
