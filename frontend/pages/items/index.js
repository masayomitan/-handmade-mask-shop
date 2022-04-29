import React, { useState, useEffect, useRef, useCallback } from "react";
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'


const ItemIndex = () => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL

  const [items, setItems] = useState([])
  const [useErrorHandler, setErrorHandler] = useState(null)

  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  useEffect(()  => {
    getItems();
  }, [])

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
      <div>
        <p>アイテム</p>
        <ul>
          {items.map((v, i) =>
            <li key={i} className="inline">
              <a href={ "/items/detail/" + v.ID}>
                  <Image
                    alt="image"
                    src={(v.ItemImages.length > 0) ? apiBaseUrl + v.ItemImages[0].file_path
                      : apiBaseUrl + "/public/img/no_image.png"}
                    width={200}
                    height={200}
                  />
                {v.Name}
              </a>
            </li>
          )}
        </ul>
      </div>
    </>
  )
}

export default ItemIndex