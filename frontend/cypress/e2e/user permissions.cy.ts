describe('Click the hyperlink for ID #2', () => {
    beforeEach(() => {
      cy.visit('/');
      cy.get('span:contains("Login")').click()
    });
  
    //This test logs in, adds a thread, adds a reply to that thread, edits the thread, deletes the thread and logs out.
    it('testing get, post, put, delete', () => {
      cy.get('input[formControlName=name]').type("test")
      cy.get('input[formControlName=password]').type("test")
      cy.get('button[color=primary]').click()

      
      cy.get('input[name=title]').type("test")
      cy.get('textarea[name=body]').type("test")
      cy.get('button[color=primary]').click()

      cy.get('p:contains("test")').click()

      cy.get('textarea[name=body]').type("test")
      cy.get('button[color=primary]').click()

      cy.get('button[aria-label="Example icon-button with a menu"]').click()
      cy.get('button[aria-label="edit button"]').click()
      cy.get('textarea[name=editedText]').type("updated test")
      cy.get('button:contains("Update Post")').click()

      
      cy.get('button[aria-label="Example icon-button with a menu"]').click()
      cy.get('button[aria-label="delete button"]').click()
      
      cy.get('span:contains("Logout")').click()
    });
  });