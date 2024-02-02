"use client";

import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import { ChangeEvent, FormEvent, useState } from "react";
import { getImage } from '../api/api';

interface ImageDisplayProps {
  loading: boolean;
  imageUrl: string | null;
}

function ImageDisplay({ loading, imageUrl }: ImageDisplayProps) {
  // ローディングインジケータを表示する
  if (loading) {
    return <div>Loading...</div>; // ここにスピナーやローディングアイコンを置く
  }

  // 画像があり、ローディングでない場合は画像を表示する
  if (imageUrl) {
    return <img src={imageUrl} alt="GitHub User Image" className="w-8/12 h-auto"/>;
  }

  // ローディングでも画像もない場合は何も表示しない
  return null;
}

function Username() {
  const [username, setUsername] = useState<string>("");
  const [imageUrl, setImageUrl] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [resultText, setResultText] = useState<string>("![GitHub persona](https://read-413014.an.r.appspot.com/create?username=");

  const copyToClipboard = async () => {
    // setResultText("![GitHub persona](https://read-413014.an.r.appspot.com/create?username=" + username + ")");
    await navigator.clipboard.writeText("![GitHub persona](https://read-413014.an.r.appspot.com/create?username=" + username + ")");
  };

  const handleSubmit = async (e: FormEvent) => {
    console.log(username);
    e.preventDefault();
    setImageUrl("http://localhost:8080/create?username=" + username);
    console.log("http://localhost:8080/create?username=" + username);
    setLoading(true);
    console.log(imageUrl);
    console.log(loading);
    setResultText("![GitHub persona](https://read-413014.an.r.appspot.com/create?username=" + username + ")");
    // const response = await fetch(`http://localhost:8080/create?username=${username}`);
    // const response = await fetch(`https://read-413014.an.r.appspot.com/create?username=${username}`);
    // const data = await response.json();
    // console.log(response);

    // const rest = await getImage(username);
    try {
      const statusCode = await getImage(username);
      console.log(loading);
      if (statusCode !== 200) {
          throw new Error(`Failed to fetch image, status code: ${statusCode}`);
      }
      console.log(loading);

    } finally {
        setLoading(false); // ローディング終了
        console.log(loading);
    }
    // console.log(rest);
    console.log(loading);

    // router.refresh();
  };

  // const copyToClipboard = (text: string) => {
  //   navigator.clipboard.writeText(text).then(() => {
  //     // alert("Copied to clipboard!");
  //   }).catch(err => {
  //     console.error('Failed to copy: ', err);
  //   });
  // }

  return (
    <form className="w-auto flex flex-col items-center justify-center mb-4 space-y-3 text-black" onSubmit={handleSubmit}>
      <input
        value={username}
        type="text"
        onChange={(e: ChangeEvent<HTMLInputElement>) =>
          setUsername(e.target.value)
        }
        className="w-64 px-4 py-2 border rounded-lg focus:outline-none focus:border-green-400"

        placeholder="Username"
      />
      <button className="w-64 px-4 py-2 text-white bg-green-500 rounded transform transition-transform duration-200 hover:bg-green-400 hover:scale-95">
        Generate
      </button>
      <div className="App">

      {imageUrl &&
      <div className="relative bg-gray-800 p-6 rounded-md">
        <div className="absolute top-1 right-1">
          <Tooltip title="Copy to Clipboard" placement="top" arrow >
            <IconButton color="primary" size="small" onClick={copyToClipboard} >
              <ContentCopyIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </div>
      <p
        className="text-white w-72 h-auto px-4 resize-none bg-transparent border-none focus:outline-none"
      >{resultText}</p>
    </div>
    }
    </div>
      {/* {imageUrl && <img src={imageUrl} alt="GitHub User Image" className=""/>} */}
      <div className="flex flex-col items-center justify-center z-50">
        <ImageDisplay loading={loading} imageUrl={imageUrl} />
      </div>
    </form>
  )
}

export default Username
