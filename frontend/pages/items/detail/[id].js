import React, { useState, useEffect, useRef, useMemo, useCallback } from "react";
import axios from 'axios';
import { useRouter } from "next/router";
import Image from 'next/image';
import Link from 'next/link';
import Header from '../../components/header.js';


const ItemDetail = () => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL

  const [item, setItem] = useState([])
  const [useErrorHandler, setErrorHandler] = useState(null)
  const router = useRouter();
  const id  = router.query.id;

  axios.defaults.baseURL = apiBaseUrl
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  

  const getDisplayItem = useCallback(async() => {
    if (id != undefined) {
      try {
        const result = await axios.get("/front/api/get-display-item/" + id);
        console.log(result.data)
        setItem(result.data)

      } catch (e) {
        if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
          console.log('400 Error!!')
          setErrorHandler(() => errorHandler(res, e))
        }
      }
    }
  }, [id])

  useEffect((id)  => {
    getDisplayItem(id)
  }, [getDisplayItem])



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
        <p>
        {item.ID}
        {item.Name}
        </p>
      </>
    )
}

export default ItemDetail