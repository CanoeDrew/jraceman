@Polymer.decorators.customElement('sql-tab')
class SqlTab extends Polymer.Element {

  @Polymer.decorators.property({type: Object, notify: true})
  queryResults: object;

  checkEnter(e: any) {
    if (e.key == 'Enter' && e.shiftKey) {
      e.stopPropagation();
      this.execute();
    }
  }

  // Clears the SQL text area.
  clear() {
    this.$.sqlText.value = "";
  }

  // Executes the SQL text.
  async execute() {
    const sql = this.$.sqlText.value;
    console.log("Execute: " + sql);     // TODO
    const path = '/api/debug/sql/';
    const formData = {
      q: sql
    };
    const options = {
      method: 'POST',
      params: formData
    };
    try {
      const result = await ApiManager.xhrJson(path, options)
      this.queryResults = result;
    } catch(e) {
      this.queryResults = {
        Error: e.responseText
      }
    }
  }
}
