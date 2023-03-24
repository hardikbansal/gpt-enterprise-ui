function AddNewButton(props){
    return (
        <button onClick={() => {props.onClick()}} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
            {props.name}
        </button>
    )
}

export  default AddNewButton;