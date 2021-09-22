import React, { useState } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

const ItemCreate = () =>  {
  const [name, setName] = useState('')

  const onSubmit = (data) => {
    
    const itemData = {
      name: data,
    };
    console.log(itemData)
    
    axios.post( '/admin/item/store', 
        itemData, {
            headers: {
              'Content-Type': 'application/json',
              'X-Requested-With': 'XMLHttpRequest'
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
    // console.log(e.target.value)
    setName(() => e.target.value)
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
          <button type="submit" onClick={() => onSubmit(name)} />登録する
      </>
  )
}

const ItemForm = document.getElementById("item_create");

ReactDOM.render(
  <ItemCreate/>, ItemForm
);

export default ItemEdit