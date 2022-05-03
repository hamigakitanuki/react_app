import React, {Component} from 'react';

export default class Todo extends Component{

  // コンポーネントの引数と初期値
  constructor(props){
    super(props);

    // vueのdata部
    this.state = {
      todos:[],
      todoName:"",
    }
  }

  onInput = (e) => {
    // 上書きする際は、連想配列を渡してあげると、キー名で指定した部分のみ更新される
    this.setState({
      todoName:e.target.value
    })
  }

  addTodo = (index) => {
    const { todos, todoName} = this.state;
    this.setState({
      todos:[...todos, todoName]
    })
  }

  removeTodo = (index) => {
    const { todos, todoName } = this.state;
    this.setState({
      todos:[...todos.slice(0, index), ...todos.slice(index + 1)]
    })
  }

  render() {
    // data部のデータをすべて取得
    const {todos} = this.state;

    return (
      <div>
        <input type="text" onInput={this.onInput}/>
        <button onClick={this.addTodo}>登録</button>
        <ul>
          {/* 波括弧で囲ってjs処理を記述 */}
          {/* プロパティに値を入れる際は、v-bindのように接頭辞をつけるわけではなく、laravelのように直接波括弧で記述 */}
          { todos.map(((todo, index) =>
            <li key="{index}">
              {todo}
              <button onClick={() => {this.removeTodo(index)}}>削除</button>
            </li>
          )) }
        </ul>
      </div>
    );
  }
}