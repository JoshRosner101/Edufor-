describe('Click the hyperlink for ID #2', () => {
    beforeEach(() => {
      cy.visit('/');
      cy.get('a:contains("The id of this thread is 2")').click()
    });
  
    it('has the correct title', () => {
      cy.title().should('equal', 'Thread #2');
    });
  });