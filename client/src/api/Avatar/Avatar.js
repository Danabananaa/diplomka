
export const uploadImage = async (file) => {
    const base64Image = await toBase64(file);
    const token = localStorage.getItem("token");
  
    try {
      const response = await fetch('/images', {
        method: "POST",
        headers: {
          Authorization: `${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ Data: base64Image }),
      });
  
      if (!response.ok) {
        const errorText = await response.text();
        console.error('Server response:', errorText);
        throw new Error('Error uploading image');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error:', error);
      throw error;
    }
  };
  
const toBase64 = (file) => // Converts image to a format that backend needs
    new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => resolve(reader.result);
      reader.onerror = (error) => reject(error);
});



export async function fetchImage(avatarName) {

    const token = localStorage.getItem("token");

    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `${token}`,
      },
    };
  
    const response = await fetch(`/images/${avatarName}`, requestOptions);
    if (!response.ok) {
      throw new Error("Failed to fetch the image");
    }
  
    const blob = await response.blob();
    const objectURL = URL.createObjectURL(blob);
  
    return objectURL;
  }