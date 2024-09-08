import { useState } from 'react';
import './App.css';

function App() {
  const [filePreviews, setFilePreviews] = useState([]);
  const [videoPreview, setVideoPreview] = useState('');

  const handleFileChange = (event) => {
    const files = Array.from(event.target.files);
    const previews = files.map(file => ({
      name: file.name,
      url: URL.createObjectURL(file)
    }));
    setFilePreviews(previews);
  };

  const handleVideoChange = async (event) => {
    const file = event.target.files[0];
    if (file) {
      const url = URL.createObjectURL(file);
      setVideoPreview(url);
    }
  };

  const handleFileUpload = async (event) => {
    event.preventDefault();
    const files = Array.from(document.getElementById('file').files);

    if (files.length === 0) {
      alert('Please select at least one file to upload.');
      return;
    }

    try {
      for (let file of files) {
        const response = await fetch('/generate-presigned-url', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ filename: file.name })
        });

        if (!response.ok) {
          throw new Error(`Failed to get presigned URL: ${response.statusText}`);
        }

        const data = await response.json();
        const presignedUrl = data.url;

        const uploadResponse = await fetch(presignedUrl, {
          method: 'PUT',
          body: file,
          headers: { 'Content-Type': file.type }
        });

        if (!uploadResponse.ok) {
          throw new Error(`Failed to upload file: ${uploadResponse.statusText}`);
        }
      }
      alert('All files uploaded successfully');
      setFilePreviews([]);
    } catch (error) {
      console.error('Error:', error);
      alert(`An error occurred while uploading the files: ${error.message}`);
    }
  };

  const handleVideoUpload = async (event) => {
    event.preventDefault();
    const file = document.getElementById('videoFile').files[0];

    if (!file) {
      alert('Please select a video to upload.');
      return;
    }

    try {
      const response = await fetch('/generate-presigned-url', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ filename: file.name })
      });

      if (!response.ok) {
        throw new Error('Failed to obtain presigned URL');
      }

      const data = await response.json();
      const presignedUrl = data.url;

      const uploadResponse = await fetch(presignedUrl, {
        method: 'PUT',
        body: file,
        headers: { 'Content-Type': file.type }
      });

      if (!uploadResponse.ok) {
        throw new Error('Failed to upload video');
      }

      alert('Video uploaded successfully');
      setVideoPreview('');
    } catch (error) {
      console.error('Error:', error);
      alert('An error occurred while uploading the video.');
    }
  };

  const handleDeleteFile = (fileToDelete) => {
    setFilePreviews(prevPreviews => prevPreviews.filter(file => file.name !== fileToDelete.name));
    const fileInput = document.getElementById('file');
    const updatedFiles = Array.from(fileInput.files).filter(file => file.name !== fileToDelete.name);
    const dataTransfer = new DataTransfer();
    updatedFiles.forEach(file => dataTransfer.items.add(file));
    fileInput.files = dataTransfer.files;
  };

  return (
    <div className="App">
      <h1 className="text-4xl font-bold text-blue-500">Upload Materials Here</h1>
      <div className="container">
        <div className="fileContainer">
          <h2 className="text-3xl font-bold text-blue-500">Upload Files</h2>
          <form id="uploadForm" onSubmit={handleFileUpload}>
            <label htmlFor="file">Choose file:</label>
            <input type="file" name="file" id="file" accept="image/*,application/pdf" multiple required onChange={handleFileChange} />
            <button className = "buttonClass" type="submit">Upload</button>
          </form>
          <h2>File Preview</h2>
          <div id="filePreviews">
            {filePreviews.map(file => (
              <div key={file.name} className="file-row">
                <img src={file.url} alt={file.name} style={{ width: '80px', height: 'auto' }} className="file-preview" />
                <span>{file.name}</span>
                <button  onClick={() => handleDeleteFile(file)} className="delete-button">Delete</button>
              </div>
            ))}
          </div>
        </div>

        <div className="wholeVideoContainer">
          <h2>Upload Video</h2>
          <form id="uploadVideoForm" onSubmit={handleVideoUpload}>
            <label htmlFor="videoFile">Choose video:</label>
            <input type="file" name="file" id="videoFile" accept="video/*" required onChange={handleVideoChange} />
            <button className = "buttonClass" type="submit">Upload</button>
          </form>
          <h2>Video Preview</h2>
          <div id="videoContainer">
            {videoPreview && (
              <video id="videoPreview" controls>
                <source src={videoPreview} type="video/mp4" />
                Your browser does not support the video tag.
              </video>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
