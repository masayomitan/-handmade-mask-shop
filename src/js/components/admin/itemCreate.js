import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

const ItemCreate = () =>  {
  const [name, setName]     = useState('')
  const [detail, setDetail]     = useState('')

  const onSubmit = (data) => {

    const itemData = {
      name: data.name,
      detail: data.detail,
    };
    console.log(itemData)
    
    axios.post( '/admin/item/store', itemData, {
            headers: {
              'Content-Type': 'application/json',
              'X-Requested-With': 'XMLHttpRequest',
              // 'X-CSRF-TOKEN' : csrf
            },
          
        }
    )
    .then(response => {
        console.log(response.data);
    })
    .catch(err => {
        console.log(err);
    });
  };

  const handleChange = (e) => {
    switch (e.target.name) {
      case 'name':
        console.log(e.target.name)
      setName(() => e.target.value)
      return
      
      case 'detail':
        console.log(e.target.name)
      setDetail(() => e.target.value)
      return
    }
  }

  return (
      <>
          <input 
            type="text"
            name="name"
            value={name}
            required="required"
            onChange={handleChange}
          />
          <input 
            type="textarea"
            name="detail"
            value={detail}
            required="required"
            onChange={handleChange}
          />

          <button type="submit" onClick={() => onSubmit(data)} />登録する
      </>
  )
}

const ItemForm = document.getElementById("item_create");

ReactDOM.render(
  <ItemCreate/>, ItemForm
);

export default ItemForm