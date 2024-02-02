export const getImage = async (username: string): Promise<number> => {
  // const response = await fetch(`http://localhost:8080/create?username=${username}`);
  const response = await fetch(`https://read-413014.an.r.appspot.com/create?username=${username}`);
  response.status;
  return response.status;
}
