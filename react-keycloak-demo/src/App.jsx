async function callApi(keycloak) {
  const res = await fetch("/api/secure", {
    headers: {
      Authorization: `Bearer ${keycloak.token}`,
    },
  });
  const text = await res.text();
  alert(text);
}

function App({keycloak}) {
  return (
    <div>
      <h1>Welcome, {keycloak.tokenParsed?.preferred_username}</h1>
      <button onClick={() => callApi(keycloak)}>Call Secure API</button>
      <button onClick={() => keycloak.logout()}>Logout</button>
    </div>
  );
}

export default App
