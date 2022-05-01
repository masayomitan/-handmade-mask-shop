import React, { useState, useEffect, useRef, useMemo, useCallback } from "react";
import axios from 'axios';
import { useRouter } from "next/router";
import Image from 'next/image';
import Link from 'next/link';
import Header from '../../components/header.js';


const ItemDetail = () => {
  const baseUrl = process.env.NEXT_PUBLIC_BASE_URL
  const apiUrl = process.env.NEXT_PUBLIC_API_URL
  const [item, setItem] = useState([])
  const [quantity, setQuantity] = useState(1)
  const [useErrorHandler, setErrorHandler] = useState(null)
  const router = useRouter();
  const id  = router.query.id;

  axios.defaults.baseURL = baseUrl
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  
  
  const getDisplayItem = useCallback(async() => {

    if (id != undefined) {
      try {
        const result = await axios.get(apiUrl + "get-display-item/" + id);
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

  const addCart = (itemId, quantity) => {
    if (!itemId) {
      return;
    }

    const request = {
      item_id: itemId,
      quantity: parseInt(quantity)
    }

    axios.post(apiUrl + 'add-order', request)
    .then(response => {
      // console.log(response)
    })
    .catch(err => {
      console.log(err);
    });
  }

  const quantityOptions = (stock) => {
    const array = []
    for (let i = 1; i < (stock + 1); i++) {
      array.push(<option key={i}>{i}</option>)
    }
    return array
  }

  const setQuantityData = (e) => {
    const t = e.target;
    const value = t.value;
    setQuantity(value)
  }

  return (
    <>
      <p>
      {item.ID}
      {item.Name}
      </p>

      <label>数量</label>
      <select
        onChange={setQuantityData}
      >
        {quantityOptions(item.Stock)}
      </select>

      <button onClick={() => addCart(item.ID, quantity)}>
        カートに追加する
      </button>
      <button>
        今すぐ購入
      </button>
    </>
  )
}

export default ItemDetail