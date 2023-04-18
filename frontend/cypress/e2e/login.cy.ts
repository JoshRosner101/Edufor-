describe('Click the hyperlink for ID #2', () => {
    beforeEach(() => {
      cy.visit('/');
      cy.get('span:contains("Login")').click()
    });
  
    it('login perms work', () => {
      //Username is incorrect
      cy.get('input[formControlName=name]').type("tes")
      cy.get('input[formControlName=password]').type("tes")
      cy.get('button[color=primary]').click()
      
      //Password is incorrect
      cy.get('input[formControlName=name]').type("t")
      cy.get('button[color=primary]').click()
      cy.wait(1000)

      //Both are correct
      cy.get('input[formControlName=password]').type("t")
      cy.get('button[color=primary]').click()
    });
  });