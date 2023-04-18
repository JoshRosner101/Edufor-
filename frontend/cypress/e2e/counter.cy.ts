describe('Click the hyperlink for ID #2', () => {
    beforeEach(() => {
      cy.visit('/');
      cy.get('mat-card-title:contains("How do I write a database in golang?")').click()
    });
  
    it('has the correct title', () => {
      cy.title().should('equal', 'Thread #2');
    });
  });