(function ($) {
    "use strict";

    // Spinner
    var spinner = function () {
        setTimeout(function () {
            if ($('#spinner').length > 0) {
                $('#spinner').removeClass('show');
            }
        }, 1);
    };
    spinner();

    // Back to top button
    $(window).scroll(function () {
        if ($(this).scrollTop() > 300) {
            $('.back-to-top').fadeIn('slow');
        } else {
            $('.back-to-top').fadeOut('slow');
        }
    });
    $('.back-to-top').click(function () {
        $('html, body').animate({scrollTop: 0}, 1500, 'easeInOutExpo');
        return false;
    });


    // Sidebar Toggler
    $('.sidebar-toggler').click(function () {
        $('.sidebar, .content').toggleClass("open");
        return false;
    });


    // Progress Bar
    $('.pg-bar').waypoint(function () {
        $('.progress .progress-bar').each(function () {
            $(this).css("width", $(this).attr("aria-valuenow") + '%');
        });
    }, {offset: '80%'});


    // Calender
    $('#calender').datetimepicker({
        inline: true,
        format: 'L'
    });


    // Testimonials carousel
    $(".testimonial-carousel").owlCarousel({
        autoplay: true,
        smartSpeed: 1000,
        items: 1,
        dots: true,
        loop: true,
        nav : false
    });


    // Chart Global Color
    Chart.defaults.color = "#6C7293";
    Chart.defaults.borderColor = "#000000";


    // Worldwide Sales Chart
    var ctx1 = $("#worldwide-sales")
    if (ctx1.length) {
        console.log(ctx1)
        ctx1 = ctx1.get(0).getContext("2d");
        var myChart1 = new Chart(ctx1, {
            type: "bar",
            data: {
                labels: ["2016", "2017", "2018", "2019", "2020", "2021", "2022"],
                datasets: [{
                        label: "USA",
                        data: [15, 30, 55, 65, 60, 80, 95],
                        backgroundColor: "rgba(235, 22, 22, .7)"
                    },
                    {
                        label: "UK",
                        data: [8, 35, 40, 60, 70, 55, 75],
                        backgroundColor: "rgba(235, 22, 22, .5)"
                    },
                    {
                        label: "AU",
                        data: [12, 25, 45, 55, 65, 70, 60],
                        backgroundColor: "rgba(235, 22, 22, .3)"
                    }
                ]
                },
            options: {
                responsive: true
            }
        });
    }


    // Salse & Revenue Chart
    var ctx2 = $("#salse-revenue")
    if (ctx2.length) {
        ctx2 = ctx2.get(0).getContext("2d");
        var myChart2 = new Chart(ctx2, {
            type: "line",
            data: {
                labels: ["2016", "2017", "2018", "2019", "2020", "2021", "2022"],
                datasets: [{
                        label: "Salse",
                        data: [15, 30, 55, 45, 70, 65, 85],
                        backgroundColor: "rgba(235, 22, 22, .7)",
                        fill: true
                    },
                    {
                        label: "Revenue",
                        data: [99, 135, 170, 130, 190, 180, 270],
                        backgroundColor: "rgba(235, 22, 22, .5)",
                        fill: true
                    }
                ]
                },
            options: {
                responsive: true
            }
        });
    }



    // Single Line Chart
    var ctx3 = $("#line-chart");
    if (ctx3.length) {
        ctx3 = ctx3.get(0).getContext("2d");
        var myChart3 = new Chart(ctx3, {
            type: "line",
            data: {
                labels: [50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150],
                datasets: [{
                    label: "Salse",
                    fill: false,
                    backgroundColor: "rgba(235, 22, 22, .7)",
                    data: [7, 8, 8, 9, 9, 9, 10, 11, 14, 14, 15]
                }]
            },
            options: {
                responsive: true
            }
        });
    }


    // Single Bar Chart
    var ctx4 = $("#bar-chart");
    if (ctx4.length) {
        ctx4 = ctx4.get(0).getContext("2d");
        var myChart4 = new Chart(ctx4, {
            type: "bar",
            data: {
                labels: ["Italy", "France", "Spain", "USA", "Argentina"],
                datasets: [{
                    backgroundColor: [
                        "rgba(235, 22, 22, .7)",
                        "rgba(235, 22, 22, .6)",
                        "rgba(235, 22, 22, .5)",
                        "rgba(235, 22, 22, .4)",
                        "rgba(235, 22, 22, .3)"
                    ],
                    data: [55, 49, 44, 24, 15]
                }]
            },
            options: {
                responsive: true
            }
        });
    }


    // Pie Chart
    var ctx5 = $("#pie-chart");
    if (ctx5.length) {
        ctx5 = ctx5.get(0).getContext("2d");
        var myChart5 = new Chart(ctx5, {
            type: "pie",
            data: {
                labels: ["Italy", "France", "Spain", "USA", "Argentina"],
                datasets: [{
                    backgroundColor: [
                        "rgba(235, 22, 22, .7)",
                        "rgba(235, 22, 22, .6)",
                        "rgba(235, 22, 22, .5)",
                        "rgba(235, 22, 22, .4)",
                        "rgba(235, 22, 22, .3)"
                    ],
                    data: [55, 49, 44, 24, 15]
                }]
            },
            options: {
                responsive: true
            }
        });
    }


    // Doughnut Chart
    var ctx6 = $("#doughnut-chart");
    if (ctx6.length) {
        ctx6 = ctx6.get(0).getContext("2d");
        var myChart6 = new Chart(ctx6, {
            type: "doughnut",
            data: {
                labels: ["Italy", "France", "Spain", "USA", "Argentina"],
                datasets: [{
                    backgroundColor: [
                        "rgba(235, 22, 22, .7)",
                        "rgba(235, 22, 22, .6)",
                        "rgba(235, 22, 22, .5)",
                        "rgba(235, 22, 22, .4)",
                        "rgba(235, 22, 22, .3)"
                    ],
                    data: [55, 49, 44, 24, 15]
                }]
            },
            options: {
                responsive: true
            }
        });
    }

    function decodeAndDecorate(chunk) {
        var elems = [];
        const decoder = new TextDecoder();
        var lines = decoder.decode(chunk).replace("\r", "\n").split("\n");
        for (var i=0; i<lines.length; i++) {
            if (lines[i] == "") {
                continue;
            }
            var elem = document.createElement("div");
            elem.className = "execution_line";
            elem.innerHTML = lines[i];
            elems.push(elem);
        }
        return elems;
    }

    // Execute buttons
    var execute_job_btn = $("#execute_job_btn");
    var execute_cancel_btn = $("#execute_cancel_btn");
    var EXECUTING_SENTINEL = false;
    //  -- cancel execution
    if (execute_cancel_btn.length) {
        execute_cancel_btn.click(function() {
            // TODO: send signal to the backend to terminate job
            EXECUTING_SENTINEL = false;
            $("#execution").html("");
            execute_cancel_btn.hide();
            execute_job_btn.show();
        })
    }
    //  -- execute
    if (execute_job_btn.length) {
        execute_cancel_btn.hide();
        execute_job_btn.click(function() {
            execute_job_btn.hide();
            fetch("http://localhost:8081/api/v1/executor_local/")
                .then(async (response) => {
                    EXECUTING_SENTINEL = true;
                    execute_cancel_btn.show();
                    $("#execution").html("");
                    const reader = response.body.getReader();
                    for await (const chunk of readChunks(reader)) {
                        var elems = decodeAndDecorate(chunk);
                        for (var i=0; i<elems.length; i++) {
                            if (EXECUTING_SENTINEL === false) {
                                $("#execution").html("");
                                break;
                            }
                            $("#execution").append(elems[i]);
                        }
                    }
                })
                .then(function() {
                    execute_job_btn.html("Re-Execute");
                    execute_job_btn.show();
                    execute_cancel_btn.hide();
                })
        });
    }

    function readChunks(reader) {
        return {
            async* [Symbol.asyncIterator]() {
                let readResult = await reader.read();
                while (!readResult.done) {
                    yield readResult.value;
                    readResult = await reader.read();
                }
            },
        };
    }


})(jQuery);
