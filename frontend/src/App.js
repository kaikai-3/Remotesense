import React, { useState } from "react";
import axios from "axios"

function App() {
    const [file, setFile] = useState(null);
    const [imageUrl, setImageUrl] =  useState("");
    const [isLoading, setIsLoading] = useState(false);

    const handleFileChange = (e) => {
        setFile(e.target.file[0]);
    }

    const handleUpload = async() => {
        if(!file) {
            alert("select an image!");
            return
        }
        const formData = new FormData();
        formData.append("file",file);

        setIsLoading(true);

        try {
            const response = await axios.post("http://localhost:8081/upload",formData,{
                headers: { "content-Type":"multipart/form-data" },
            });
            setImageUrl("http://localhost:8081"+response.data.url);
        }catch(error){
            console.error("Failed to upload:",error);
            alert("Failed to upload!");
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div style={{ textAlign: "center",marginTop: "50px"}}>
            <h1>图片上传与处理</h1>
            <input type="file" onChange={handleFileChange}/>
            <button onClick={handleUpload} disabled={isLoading}>
                {isLoading ? "处理中...":"上传"}
            </button>
            {imageUrl &&(
                <div style={{ marginTop: "20px"}}>
                    <h2>处理后的图片：</h2>
                    <img src={imageUrl} alt="processed" style={{maxWidth: "100%"}}/>
                </div>
            )}
        </div>
    );
}

export default App;