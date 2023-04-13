import React, { useState } from 'react';
import { uploadImage } from '../../api/Avatar/Avatar';

const ImageUploadForm = () => { // IMAGE Upload form on profile page
  const [file, setFile] = useState(null);
  const [errorMessage, setErrorMessage] = useState('');

  const handleImageChange = (e) => {
    const selectedFile = e.target.files[0];
    const allowedExtensions = ['image/jpeg', 'image/png', 'image/jpg'];
    const fileSizeLimit = 5 * 1024 * 1024; // 5 MB

    if (selectedFile && allowedExtensions.includes(selectedFile.type) && selectedFile.size <= fileSizeLimit) {
      setFile(selectedFile);
      setErrorMessage('');
    } else {
      setFile(null);
      setErrorMessage('Please upload a png, jpg, or jpeg image with a size of up to 5 MB.');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (file) {
      await uploadImage(file);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input type="file" onChange={handleImageChange} />
        <button type="submit">Submit</button>
      </form>
      {errorMessage && <p>{errorMessage}</p>}
    </div>
  );
};

export default ImageUploadForm;