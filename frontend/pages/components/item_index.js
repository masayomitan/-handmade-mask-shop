import React, { useState, useEffect, useRef, useCallback } from "react";
import axios from 'axios';
import Image from 'next/image'
import Link from 'next/link';

const ItemIndex = () => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL

  const [items, setItems] = useState([])
  const [useErrorHandler, setErrorHandler] = useState(null)
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  const getItems = useCallback(async () => {
    try {
      const result = await axios.get("front/api/get-display-items");
      setItems(result.data)
    } catch (e) {
      if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
        console.log('400 Error!!')
        setErrorHandler(() => errorHandler(res, e))
      }
      
    }
  }, [])
  
  useEffect(()  => {
    getItems();
  }, [getItems])

  const errorHandler = (res, error) => {
    if (error.response) {
        res.status(error.response.status).send({
            error: error.response.data,
            errorMsg: error.message
        })
    } else {
        res.status(500).send({ errorMsg: error.message })
    }
  }
  
  return (
    <>
      {items.map((v, i) =>
          <li key={i} className="inline">
            <a href={ "/items/detail/" + v.ID}>
              {v.ItemImages.length > 0 ?
                <Image
                  alt="image"
                  src={apiBaseUrl + v.ItemImages[0].file_path}
                  width={300}
                  height={300}
                />
                :   
                <Image
                  alt="image"
                  src={apiBaseUrl + "/public/img/no_image.png"}
                  width={300}
                  height={300}
                />
              } 
              {v.Name}
            </a>
          </li>
        )}
    </>
  )
}

export default ItemIndex