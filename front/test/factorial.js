process.env.NODE_ENV = 'test';

//Require the dev-dependencies
let chai = require('chai');
let chaiHttp = require('chai-http');
let should = chai.should();

chai.use(chaiHttp);
//Our parent block
describe('Factorial', () => {

 /*
  * Test the /GET route
  */
  describe('/GET 4', () => {
      it('it should calculate factorial of 4', (done) => {
        chai.request("http://localhost:8080")
            .get('/4')
            .end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.Factorial.should.equal("24");
              done();
              process.exit(1);
            });
      });
  });

});