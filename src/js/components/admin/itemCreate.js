import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

const ItemCreate = () =>  {
  const [data, setData]     = useState([])

  const onSubmit = (data) => {
    
    const itemData = {
      name: data.name,
      detail: data.detail,
    };
    
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
      const name = e.target.name;
      const value = e.target.value;
      setData({...data, [name]: value})
  }

  return (
      <>
          <input 
            type="text"
            name="name"
            // value={data.name}
            required="required"
            onChange={handleChange}
          />
          <input 
            type="textarea"
            name="detail"
            // value={data.detail}
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