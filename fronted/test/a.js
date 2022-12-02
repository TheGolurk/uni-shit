var modalWrap = null;
var modalWrapErr = null;

const showOKMessage = (
    title,
    description,
    ) =>  {
    if (modalWrap !== null) {
        modalWrap.remove();
    }

    modalWrap = document.createElement('div');
    modalWrap.innerHTML = `
    <div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
        aria-labelledby="staticBackdropLabel" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-7" id="staticBackdropLabel">${title}</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    ${description}
                </div>
                <div class="modal-footer">
                    <button data-bs-dismiss="modal" type="button" class="btn btn-primary"
                            style="background-color: forestgreen;">CONTINUAR</button>
                </div>
            </div>
        </div>
    </div>
    
    `

    document.body.append(modalWrap);
    var modal = new bootstrap.Modal(modalWrap.querySelector('.modal'));
    modal.show();
}

const showErrMessage = (
    title,
    description,
    ) =>  {
    if (modalWrapErr !== null) {
        modalWrapErr.remove();
    }

    modalWrapErr = document.createElement('div');
    modalWrapErr.innerHTML = `
    <div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
    aria-labelledby="staticBackdropLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header">
                <h1 class="modal-title fs-7" id="staticBackdropLabel">${title}</h1>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                ${description}
            </div>
            <div class="modal-footer">
                <button data-bs-dismiss="modal" type="button" class="btn btn-primary"
                        style="background-color: crimson;">CONTINUAR</button>
            </div>
        </div>
    </div>
</div>
    
    `

    document.body.append(modalWrapErr);
    var modal = new bootstrap.Modal(modalWrapErr.querySelector('.modal'));
    modal.show();
}